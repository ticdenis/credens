.PHONY: docker-build docker-shell go-install go-build-app go-run-app go-build-run-app go-test-app go-test-lib

# Variables

DOCKER_IMAGE = credens/go
DOCKER_VOLUME = credens_go_volume
DOCKER_CONFIG_DIR = ./docker/

GO_CGO_ENABLED = 0
GO_BUILD_APPS_DIR = ./builds/apps/
GO_APPS_DIR = ./apps/
GO_APPS_TESTS_DIR = ./tests/apps/
GO_LIBS_DIR = ./libs/
GO_LIBS_TESTS_DIR = ./tests/libs/

# Scripts

docker-build:
	@docker build --file ${DOCKER_CONFIG_DIR}local/go/Dockerfile -t ${DOCKER_IMAGE} .

docker-shell:
	@docker run --rm --name=${DOCKER_VOLUME} -v ${DOCKER_VOLUME}:/go/src -it ${DOCKER_IMAGE} bash

go-install:
	@go mod download

go-build-app:
	@CGO_ENABLED=${GO_CGO_ENABLED} go build -o ${GO_BUILD_APPS_DIR}${name} ${GO_APPS_DIR}${name}

go-run-app:
	@${GO_BUILD_APPS_DIR}${name}

go-build-run-app: go-build-app go-run-app

go-test-app:
	@CGO_ENABLED=${GO_CGO_ENABLED} go test ${GO_APPS_TESTS_DIR}${name}/*

go-test-lib:
	@CGO_ENABLED=${GO_CGO_ENABLED} go test ${GO_LIBS_TESTS_DIR}${name}/*
