#!/bin/bash

set -x
set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT=$(cd $(dirname "$0")/.. ; pwd)
VERSION="v1alpha1"

go install -mod=mod sigs.k8s.io/controller-tools/cmd/controller-gen

go build -o dist/openapi-gen ./vendor/k8s.io/kube-openapi/cmd/openapi-gen
./dist/openapi-gen \
  --go-header-file ${PROJECT_ROOT}/hack/custom-boilerplate.go.txt \
  --input-dirs github.com/argoproj/argo-cd/pkg/apis/application/${VERSION} \
  --output-package github.com/argoproj/argo-cd/pkg/apis/application/${VERSION} \
  --report-filename pkg/apis/api-rules/violation_exceptions.list \
  $@

which controller-gen
controller-gen paths=./pkg/apis/application/... crd:trivialVersions=true output:crd:stdout

go build -o ./dist/gen-crd-spec ${PROJECT_ROOT}/hack/gen-crd-spec
./dist/gen-crd-spec
