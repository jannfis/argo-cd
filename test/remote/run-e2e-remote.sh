#!/bin/sh

# Should point to the Argo CD API endpoint on the cluster
if test "${ARGOCD_SERVER}" = ""; then
	echo "Please set ARGOCD_SERVER to the remote Argo CD API endpoint to test." >&2
	exit 1
fi

# ARGOCD_E2E_REMOTE must be set to 'true' in order for remote tests to work
export ARGOCD_E2E_REMOTE=true

# Set ARGOCD_E2E_NAMESPACE to the namespace the Argo CD we're testing against is
# running in. Defaults to "argocd-e2e"
export ARGOCD_E2E_NAMESPACE=

# Name prefix the operator sets on resources created for Argo CD instance. This
# is usually also the name of the instance itself.
export ARGOCD_E2E_NAME_PREFIX="argocd-test"

# When our 
export ARGOCD_E2E_DYNAMIC_PERMISSIONS=false

# This is to skip some (deprecated) tests
export ARGOCD_E2E_K3S=true

# Skip some tests
# GnuPG features not yet available with GitOps Operator
export ARGOCD_E2E_SKIP_GPG=true
# Some tests do not work OOTB with OpenShift
export ARGOCD_E2E_SKIP_OPENSHIFT=true
# Skip Helm tests
export ARGOCD_E2E_SKIP_HELM=true

# Unauthenticated URLs for pushing from CI
#
# Use `kubectl port-forward service/argocd-e2e-server 9081:9081` to set up the
# listener required for this.
export ARGOCD_E2E_GIT_SERVICE="http://127.0.0.1:9081/argo-e2e/testdata.git"
export ARGOCD_E2E_HELM_SERVICE="http://127.0.0.1:9081/helm-repo"
export ARGOCD_E2E_GIT_SERVICE_SUBMODULE="http://127.0.0.1:9081/argo-e2e/submodule.git"
export ARGOCD_E2E_GIT_SERVICE_SUBMODULE_PARENT="http://127.0.0.1:9081/argo-e2e/submoduleParent.git"

# URLs used during testing - usually no need to change thos
export ARGOCD_E2E_REPO_SSH="ssh://root@argocd-e2e-server:2222/tmp/argo-e2e/testdata.git"
export ARGOCD_E2E_REPO_SSH_SUBMODULE="ssh://root@argocd-e2e-server:2222/tmp/argo-e2e/submodule.git"
export ARGOCD_E2E_REPO_SSH_SUBMODULE_PARENT="ssh://root@argocd-e2e-server:2222/tmp/argo-e2e/submoduleParent.git"
export ARGOCD_E2E_REPO_HTTPS="https://argocd-e2e-server:9443/argo-e2e/testdata.git"
export ARGOCD_E2E_REPO_HTTPS_CLIENT_CERT="https://argocd-e2e-server:9444/argo-e2e/testdata.git"
export ARGOCD_E2E_REPO_HTTPS_SUBMODULE="https://argocd-e2e-server:9443/argo-e2e/submodule.git"
export ARGOCD_E2E_REPO_HTTPS_SUBMODULE_PARENT="https://argocd-e2e-server:9443/argo-e2e/submoduleParent.git"
export ARGOCD_E2E_REPO_HELM="https://argocd-e2e-server:9444/helm-repo"
export ARGOCD_E2E_REPO_DEFAULT="http://argocd-e2e-server:9081/argo-e2e/testdata.git"

$*
