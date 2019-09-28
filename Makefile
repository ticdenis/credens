.PHONY: help available-apps available-libs build-docker shell install-deps install-dep build-app run-app build-run-app test-app test-lib

DOCKER_IMAGE = credens/go
DOCKER_VOLUME = credens_go_vol
GO_ENVS = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

help: ## It displays this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

available-apps: ## It shows available project's apps.
	@ls ./apps

available-libs: ## It shows available project's libs.
	@ls ./libs

build-docker: ## It builds a Docker image for this project.
	@docker build -t ${DOCKER_IMAGE} .

shell: ## It executes a shell inside a Docker container with optional "$args" arg.
	@docker run --rm --name=${DOCKER_VOLUME} -v ${PWD}:/app -it ${args} ${DOCKER_IMAGE} bash

install-deps: ## It install go dependencies with go.mod file.
	@go mod vendor && go mod download

install-dep:  ## It install go dependency with "$pkg" arg.
	@go get -t ${pkg} && go mod vendor && go mod download || \
	echo "'pkg' argument is required to install a dependency"

build-app: ## It builds app given with "$name" arg.
	@${GO_ENVS} go build -mod=vendor -o ./builds/apps ./apps/${name} || \
    echo "'name' argument is required to build an app"

run-app: ## It run app given with "$name" arg.
	@./builds/apps/${name} || \
	echo "'name' argument is required to run an app if exists"

build-run-app: build-app run-app ## It builds and runs app given with "$name" arg.

test-app: ## It runs app tests given with "$name" arg.
	@${GO_ENVS} go test ./tests/apps/${name}/* || \
    echo "'name' argument is required to run app tests if they exists"

test-lib: ## It runs lib tests given with "$name" arg.
	@${GO_ENVS} go test ./tests/libs/${name}/* || \
    echo "'name' argument is required to run lib tests if they exists"
