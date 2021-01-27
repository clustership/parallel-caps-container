NS ?= xymox
VERSION ?= 0.1.1
REPOSITORY ?= quay.io

IMAGE_NAME ?= devsecops-parallel
DOCKER = docker

DOCKERFILE = Dockerfile.golang

.PHONY: build push

build-nocache: Dockerfile
	$(DOCKER) build --rm --no-cache -t $(NS)/$(IMAGE_NAME):$(VERSION) -f $(DOCKERFILE) .
build:
	$(DOCKER) build -t $(NS)/$(IMAGE_NAME):$(VERSION) -f $(DOCKERFILE) .
	$(DOCKER) tag $(NS)/$(IMAGE_NAME):$(VERSION) $(REPOSITORY)/$(NS)/$(IMAGE_NAME):$(VERSION)

tag:
	$(DOCKER) tag $(NS)/$(IMAGE_NAME):$(VERSION) $(REPOSITORY)/$(NS)/$(IMAGE_NAME):$(VERSION)

push: build
	$(DOCKER) tag $(NS)/$(IMAGE_NAME):$(VERSION) $(REPOSITORY)/$(NS)/$(IMAGE_NAME):latest
	$(DOCKER) push $(REPOSITORY)/$(NS)/$(IMAGE_NAME):$(VERSION)
	$(DOCKER) push $(REPOSITORY)/$(NS)/$(IMAGE_NAME):latest

shell:
	$(DOCKER) run --rm --name $(IMAGE_NAME) -ti $(REPOSITORY)/$(NS)/$(IMAGE_NAME):$(VERSION) /bin/bash

tag-latest: tag
	$(DOCKER) tag $(NS)/$(IMAGE_NAME):$(VERSION) $(REPOSITORY)/$(NS)/$(IMAGE_NAME):latest
  
push-latest: tag-latest
  make push -e VERSION=latest

release: push-latest push tag build

default: build
