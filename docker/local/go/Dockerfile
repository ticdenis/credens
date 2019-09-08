ARG GOLANG_VERSION=1.13-alpine

FROM golang:$GOLANG_VERSION

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make

WORKDIR /go/src

COPY go.mod go.sum vendor ./

RUN go mod download

COPY . .

RUN make build-cli
RUN make build-http

CMD bash