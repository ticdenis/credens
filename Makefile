IMAGE = credens/go

build-docker:
	@docker build --file ./docker/local/go/Dockerfile -t ${IMAGE} .

build-go:
	@go build ./...

shell:
	@docker run --rm -it ${IMAGE} bash