ARG GOLANG_VERSION=1.13-alpine

FROM golang:$GOLANG_VERSION

ENV CGO_ENABLED=0

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make

WORKDIR /go/src

COPY go.mod go.sum vendor ./

RUN go mod download

COPY . .

RUN go build -o ./...

CMD go run ./src/user_interface/cli/main.go