#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
. ${SCRIPT_ROOT}/hack/versions.sh
TARGET_SCRIPT=${SCRIPT_ROOT}/dist/generate-groups.sh
CODEGEN_PKG=$(go list -mod=readonly -m -f "{{.Dir}}" k8s.io/code-generator)
sed -e '/go install .\/cmd/d' ${CODEGEN_PKG}/generate-groups.sh > ${TARGET_SCRIPT}
chmod +x ${TARGET_SCRIPT}

go install ${CODEGEN_PKG}/cmd/{defaulter-gen,client-gen,lister-gen,informer-gen,deepcopy-gen}

GOPATH="$HOME/go" bash -x ${CODEGEN_PKG}/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/argoproj/argo-cd/pkg/client github.com/argoproj/argo-cd/pkg/apis \
  "application:v1alpha1" \
  --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt
