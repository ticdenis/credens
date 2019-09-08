IMAGE = credens/go

.PHONY: build-cli run-cli build-http run-http

build-docker:
	@docker build --file ./docker/local/go/Dockerfile --rm -t ${IMAGE} .

install:
	@go mod download

build-cli:
	@CGO_ENABLED=0 go build -o ./var/build/cli ./src/user_interface/cli

run-cli:
	@var/build/cli

cli: build-cli run-cli

build-http:
	@CGO_ENABLED=0 go build -o ./var/build/http ./src/user_interface/http

run-http:
	@var/build/http

http: build-http run-http

shell:
	@docker run --rm -it ${IMAGE} bash