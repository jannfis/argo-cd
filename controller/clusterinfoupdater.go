package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/argoproj/gitops-engine/pkg/cache"
	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"github.com/argoproj/gitops-engine/pkg/utils/tracing"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	statecache "github.com/argoproj/argo-cd/v2/controller/cache"
	appv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/pkg/client/listers/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/util/argo"
	appstatecache "github.com/argoproj/argo-cd/v2/util/cache/appstate"
	"github.com/argoproj/argo-cd/v2/util/db"
)

const (
	secretUpdateInterval = 10 * time.Second
	pingInterval         = 10 * time.Second // TODO: Make this interval configurable and provide reasonable default
)

type clusterInfoUpdater struct {
	infoSource    statecache.LiveStateCache
	db            db.ArgoDB
	appLister     v1alpha1.ApplicationNamespaceLister
	cache         *appstatecache.Cache
	clusterFilter func(cluster *appv1.Cluster) bool
	projGetter    func(app *appv1.Application) (*appv1.AppProject, error)
	namespace     string
}

func NewClusterInfoUpdater(
	infoSource statecache.LiveStateCache,
	db db.ArgoDB,
	appLister v1alpha1.ApplicationNamespaceLister,
	cache *appstatecache.Cache,
	clusterFilter func(cluster *appv1.Cluster) bool,
	projGetter func(app *appv1.Application) (*appv1.AppProject, error),
	namespace string) *clusterInfoUpdater {

	return &clusterInfoUpdater{infoSource, db, appLister, cache, clusterFilter, projGetter, namespace}
}

func (c *clusterInfoUpdater) Run(ctx context.Context) {
	c.updateClusters()
	updateTicker := time.NewTicker(secretUpdateInterval)
	pingTicker := time.NewTicker(pingInterval)
	pingChan := make(chan bool)
OUTER:
	for {
		select {
		case <-ctx.Done():
			updateTicker.Stop()
			pingTicker.Stop()
			break OUTER
		case <-updateTicker.C:
			c.updateClusters()
		case <-pingTicker.C:
			// Pinging the clusters may take some time. We stop the ticker,
			// so that not multiple pings will run at the same time, and we
			// reset the ticker once the ping has completed.
			pingTicker.Stop()
			// We spawn the ping process in its own goroutine, so it does
			// not interfere with the cluster update interval.
			go c.pingClusters(pingChan)
		case <-pingChan:
			pingTicker.Reset(pingInterval)
		}
	}
}

// pingClusters sends a version request to remote clusters' APIs. If this
// fails, it will invalidate the corresponding cluster cache so the cluster
// will be marked as failed, if the error does not recover meanwhile.
func (c *clusterInfoUpdater) pingClusters(ch chan bool) {
	defer func() {
		ch <- true
	}()
	clusters, err := c.getClustersFiltered()
	if err != nil {
		log.Errorf("Could not retrieve list of clusters for cluster ping: %v", err)
		return
	}
	_ = kube.RunAllAsync(len(clusters), func(i int) error {
		for _, ci := range c.infoSource.GetClustersInfo() {
			if ci.Server == clusters[i].Server {
				c.pingCluster(clusters[i], ci)
				return nil
			}
		}
		return nil
	})
}

// pingCluster will perform a ping-style request to the given cluster, and will
// execute a cluster cache invalidation on a failing cluster, and a re-sync on
// a recovering one.
func (c *clusterInfoUpdater) pingCluster(cluster appv1.Cluster, info cache.ClusterInfo) {
	logCtx := log.WithField("server", cluster.Server)
	logCtx.Debugf("Pinging cluster to check availability")

	k := kube.KubectlCmd{Tracer: tracing.NopTracer{}}
	// TODO: Do we want configurable retry for the version request?
	// TODO: Do we want configurable timeout for the version reqest?
	v, err := k.GetServerVersion(cluster.RESTConfig())
	if err != nil {
		// Ping error and last sync error is nil means we just lost the cluster
		if info.SyncError == nil {
			logCtx.Infof("Cluster became unavailable. Reason: '%v')", err)
			if err := c.invalidateClusterCache(cluster); err != nil {
				logCtx.Infof("Failed to invalidate cluster cache: %v", err)
				return
			}
		} else {
			logCtx.Debugf("Cluster still unavailable")
		}
	} else {
		// Ping success and last sync error is set means cluster just recovered
		if info.SyncError != nil {
			logCtx.Infof("Cluster became available again after %vs", time.Since(*info.LastCacheSyncTime).Seconds())
			if err := c.invalidateClusterCache(cluster); err != nil {
				logCtx.Info("Failed to mark cluster as available: %v", err)
			}
		} else {
			logCtx.Debugf("Cluster available. Version: '%v'", v)
		}
	}

	c.updateClusterInfo(cluster, &info)
}

// invalidateClusterCache invalidates a cluster cache and subsequentially
// ensures it is synced again.
func (c *clusterInfoUpdater) invalidateClusterCache(cluster appv1.Cluster) error {
	cache, err := c.infoSource.GetClusterCache(cluster.Server)
	if err != nil {
		return err
	}
	cache.Invalidate()
	err = cache.EnsureSynced()
	return err
}

func (c *clusterInfoUpdater) updateClusters() {
	infoByServer := make(map[string]*cache.ClusterInfo)
	clustersInfo := c.infoSource.GetClustersInfo()
	for i := range clustersInfo {
		info := clustersInfo[i]
		infoByServer[info.Server] = &info
	}
	clusters, err := c.getClustersFiltered()
	if err != nil {
		log.Warnf("Failed to save clusters info: %v", err)
		return
	}
	_ = kube.RunAllAsync(len(clusters), func(i int) error {
		cluster := clusters[i]
		if err := c.updateClusterInfo(cluster, infoByServer[cluster.Server]); err != nil {
			log.Warnf("Failed to save clusters info: %v", err)
		}
		return nil
	})
	log.Debugf("Successfully saved info of %d clusters", len(clusters))
}

func (c *clusterInfoUpdater) updateClusterInfo(cluster appv1.Cluster, info *cache.ClusterInfo) error {
	apps, err := c.appLister.List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error while fetching the apps list: %w", err)
	}
	var appCount int64
	for _, a := range apps {
		if c.projGetter != nil {
			proj, err := c.projGetter(a)
			if err != nil || !proj.IsAppNamespacePermitted(a, c.namespace) {
				continue
			}
		}
		if err := argo.ValidateDestination(context.Background(), &a.Spec.Destination, c.db); err != nil {
			continue
		}
		if a.Spec.Destination.Server == cluster.Server {
			appCount += 1
		}
	}
	now := metav1.Now()
	clusterInfo := appv1.ClusterInfo{
		ConnectionState:   appv1.ConnectionState{ModifiedAt: &now},
		ApplicationsCount: appCount,
	}
	if info != nil {
		clusterInfo.ServerVersion = info.K8SVersion
		clusterInfo.APIVersions = argo.APIResourcesToStrings(info.APIResources, true)
		if info.LastCacheSyncTime == nil {
			clusterInfo.ConnectionState.Status = appv1.ConnectionStatusUnknown
		} else if info.SyncError == nil {
			clusterInfo.ConnectionState.Status = appv1.ConnectionStatusSuccessful
			syncTime := metav1.NewTime(*info.LastCacheSyncTime)
			clusterInfo.CacheInfo.LastCacheSyncTime = &syncTime
			clusterInfo.CacheInfo.APIsCount = int64(info.APIsCount)
			clusterInfo.CacheInfo.ResourcesCount = int64(info.ResourcesCount)
		} else {
			clusterInfo.ConnectionState.Status = appv1.ConnectionStatusFailed
			clusterInfo.ConnectionState.Message = info.SyncError.Error()
		}
	} else {
		clusterInfo.ConnectionState.Status = appv1.ConnectionStatusUnknown
		if appCount == 0 {
			clusterInfo.ConnectionState.Message = "Cluster has no applications and is not being monitored."
		}
	}

	return c.cache.SetClusterInfo(cluster.Server, &clusterInfo)
}

// getClustersFiltered returns a list of clusters from the configuration that
// this controller is responsible for.
func (c *clusterInfoUpdater) getClustersFiltered() ([]appv1.Cluster, error) {
	clusters, err := c.db.ListClusters(context.Background())
	if err != nil {
		return nil, err
	}
	var clustersFiltered []appv1.Cluster
	if c.clusterFilter == nil {
		clustersFiltered = clusters.Items
	} else {
		for i := range clusters.Items {
			if c.clusterFilter(&clusters.Items[i]) {
				clustersFiltered = append(clustersFiltered, clusters.Items[i])
			}
		}
	}
	return clustersFiltered, nil
}
