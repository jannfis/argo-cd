# End-to-end tests against a real cluster

Using the tools in this directory, you can run the End-to-End testsuite against
a real Argo CD workload, that is deployed to a K8s cluster, instead of running
it against a locally running Argo CD.

Since e2e tests are destructive, do **not** run it against an installation that
you depend on.

## Preparations

### Build the repository container image

You will need to build & publish a container image that will hold the required
testing repositories.

This container image will be named `argocd-e2e-cluster`, so you will need to
setup a corresponding repository for it in your registry as well.

To build it, run the following. Note that kustomize is required:

```shell
export IMAGE_NAMESPACE=quay.io/youruser
# builds & tags the image
make image
# pushes the image to your repository
make image-push
# rebuild the manifests
make manifests
```

If you do not have kustomize installed, you need to manually edit the manifests
at `test/remote/manifests/e2e_repositories.yaml` to point to the correct image.

### Install the Argo CD you want to test

It is recommended to install in the `argocd-e2e` namespace:

```
kubectl create ns argocd-e2e
kubectl -n argocd-e2e apply -f <your Argo CD installation manifests>
```

### 