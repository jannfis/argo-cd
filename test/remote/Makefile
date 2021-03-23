PWD=$(shell pwd)
TEST_ROOT=$(shell realpath $(PWD)/../..)

IMAGE_NAMESPACE?=quay.io/jfischer-redhat
IMAGE_NAME=argocd-e2e-cluster
IMAGE_TAG=latest
ifneq (IMAGE_NAMESPACE, '')
IMAGE_PREFIX=$(IMAGE_NAMESPACE)/
else
IMAGE_PREFIX=
endif

image:
	docker build -t $(IMAGE_PREFIX)$(IMAGE_NAME):$(IMAGE_TAG) -f Dockerfile $(TEST_ROOT)

image-push: image
	docker push $(IMAGE_PREFIX)$(IMAGE_NAME):$(IMAGE_TAG)
