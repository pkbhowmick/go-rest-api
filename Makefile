.PHONY: all build-binary build-image push clean install uninstall

all: build-binary build-image push


TAG ?= 2.0.1
REGISTRY ?= registry.hub.docker.com
APP_NAME ?= go-rest-api
DOCKER_REPO ?= pkbhowmick
RELEASE_NAME ?= go-api-server


build-binary:
	@echo Building Go binary file
	go build -o ${APP_NAME} .


build-image: build-binary
	@echo Building the API Server Project ...
	docker build -t ${DOCKER_REPO}/${APP_NAME}:${TAG} .


push: build-image
	@echo Pushing the Image into ${REGISTRY}
	docker push ${DOCKER_REPO}/${APP_NAME}:${TAG}


install:
	helm install ${RELEASE_NAME} chart

uninstall:
	helm uninstall ${RELEASE_NAME}

clean:
	@echo Cleaning up...
	rm -rf ${APP_NAME}
