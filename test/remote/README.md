# End-to-end tests against a real cluster

Using the tools in this directory, you can run the End-to-End testsuite against
a real Argo CD workload, that is deployed to a K8s cluster, instead of running
it against a locally running Argo CD.

Since e2e tests are destructive, do **not** run it against an installation that
you depend on.

## Preparations

### Install the Argo CD you want to test

It is recommended to install in the `argocd-e2e` namespace:

```shell
kubectl create ns argocd-e2e
kubectl -n argocd-e2e apply -f <your Argo CD installation manifests>
```

### Give the Argo CD the appropriate RBAC permissions

```shell
# If you installed to a different namespace, set accordingly
export NAMESPACE=argocd-e2e
# If you installed Argo CD via Operator, set accordingly
# export ARGOCD_E2E_NAME_PREFIX=argocd-cluster
./test/remote/generate-permissions.sh | kubectl apply -f -
```

### Build the repository container image

You will need to build & publish a container image that will hold the required
testing repositories.

This container image will be named `argocd-e2e-cluster`, so you will need to
setup a corresponding repository for it in your registry as well.

To build it, run the following. Note that kustomize is required:

```shell
cd test/remote
export IMAGE_NAMESPACE=quay.io/youruser
# builds & tags the image
make image
# pushes the image to your repository
make image-push
# build the manifests & store them at temp location
make manifests > /tmp/e2e-repositories.yaml
```

If you do not have kustomize installed, you need to manually edit the manifests
at `test/remote/manifests/e2e_repositories.yaml` to point to the correct image.

### Deploy the test container and additional permissions

**Note:** The test container requires to be run in privileged mode for now, due
to some processes running as root (this may change some day...).

```shell
kubectl -n argocd-e2e apply -f /tmp/e2e-repositories.yaml
```

Verify that the deployment was succesful:

```shell
kubectl -n argocd-e2e rollout status deployment argocd-e2e-cluster
```

## Start the tests

### On local cluster (e.g. K3s, microk8s, minishift)

Set the server endpoint of the Argo CD API. If you are running on the same host
as the cluster, or the cluster IPs are routed to your host, you can use the
following:

```shell
export ARGOCD_SERVER=$(kubectl get svc argocd-server -o jsonpath='{.spec.clusterIP}')
```

Set the admin password to use:

```shell
export ARGOCD_E2E_ADMIN_PASSWORD=$(kubectl get secrets argocd-initial-admin-secret -o jsonpath='{.data.password}'|base64 -d)
```

### On remote OpenShift cluster


