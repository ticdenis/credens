ARG GOLANG_VERSION=1.13-alpine

FROM golang:$GOLANG_VERSION

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make

ENV GO111MODULE=on

WORKDIR /app

COPY . .

CMD bash
