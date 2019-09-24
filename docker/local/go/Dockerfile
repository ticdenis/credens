ARG GOLANG_VERSION=1.13-alpine

FROM golang:$GOLANG_VERSION

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make

WORKDIR /go/src

COPY Makefile go.mod go.sum ./

RUN make go-install

COPY . .

#RUN make go-build-app name=cli
#RUN make go-build-app name=grpc
#RUN make go-build-app name=http
#RUN make go-build-app name=rpc

CMD bash
