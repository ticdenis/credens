.PHONY: help available-apps available-libs build-docker shell install-deps install-dep build-app run-app build-run-app test-app coverage-app test-lib coverage-lib

DOCKER_IMAGE = credens/go
DOCKER_VOLUME = credens_go_vol
DOCKER_NETWORK = credens_go_network
GO_ENVS = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

help: ## It displays this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

available-apps: ## It shows available project's apps.
	@ls ./apps

available-libs: ## It shows available project's libs.
	@ls ./libs

up-services: ## It ups Docker services defined at docker-compose.yml file.
	@docker-compose up -d

down-services: ## It downs Docker services defined at docker-compose.yml file.
	@docker-compose down

build-docker: ## It builds a Docker image for this project.
	@mkdir -p var/builds/apps && mkdir -p var/dockers/db
	@docker network inspect ${DOCKER_NETWORK} &>/dev/null || docker network create ${DOCKER_NETWORK}
	@docker build -t ${DOCKER_IMAGE} .

shell: ## It executes a shell inside a Docker container with optional "$args" arg.
	@docker run --rm --network ${DOCKER_NETWORK} --name=${DOCKER_VOLUME} -v ${PWD}:/app -it ${args} ${DOCKER_IMAGE} bash

shell-service: ## It executes a shell inside a Docker service container with optional "$name" arg.
	@docker exec -it ${name} bash

install-deps: ## It install go dependencies with go.mod file.
	@go mod vendor && go mod download

install-dep:  ## It install go dependency with "$pkg" arg.
	@go get -t ${pkg} && go mod vendor && go mod download || \
	echo "'pkg' argument is required to install a dependency"

build-app: ## It builds app given with "$name" arg.
	@${GO_ENVS} go build -mod=vendor -o ./var/builds/apps ./apps/${name} || \
    echo "'name' argument is required to build an app"

run-app: ## It run app given with "$name" arg and "$args" optional arg.
	@./var/builds/apps/${name} ${args} || \
	echo "'name' argument is required to run an app if exists"

build-run-app: build-app run-app ## It builds and runs app given with "$name" arg and "$args" optional arg.

test-app: ## It runs app tests given with "$name" arg.
	@${GO_ENVS} go test ./tests/apps/${name}/* || \
    echo "'name' argument is required to run app tests if they exists"

coverage-app: ## It runs app tests with coverage given with "$name" arg.
	@${GO_ENVS} go test ./tests/apps/${name}/* -cover || \
    echo "'name' argument is required to run app tests with coverage if they exists"

test-lib: ## It runs lib tests given with "$name" arg.
	@${GO_ENVS} go test ./tests/libs/${name}/* || \
    echo "'name' argument is required to run lib tests if they exists"

coverage-lib: ## It runs lib tests with coverage given with "$name" arg.
	@${GO_ENVS} go test ./tests/libs/${name}/* -cover || \
    echo "'name' argument is required to run lib tests with coverage if they exists"