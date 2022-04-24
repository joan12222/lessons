.DEFAULT_GOAL := docker-image

IMAGE ?= flyinpenguin/demo8-server:metrics

image/demo8-server: $(shell find . -name '*.go')
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $@ ./main.go

.PHONY: docker-image
docker-image: image/demo8-server
	docker build -t $(IMAGE) image/

.PHONY: push-image
push-image: docker-image
	docker push $(IMAGE)
