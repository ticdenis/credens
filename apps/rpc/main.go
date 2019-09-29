package main

import (
	"credens/apps/rpc/config"
	"credens/apps/rpc/service"
	"credens/libs/shared/infrastructure/logging"
	"fmt"
	"github.com/defval/inject"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	env, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	container, err := config.BuildContainer(*env)
	if err != nil {
		panic(err)
	}

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func run(container *inject.Container, env config.Environment) error {
	var logger logging.Logger
	if err := container.Extract(&logger); err != nil {
		return err
	}

	rcvr := new(service.RPCAPIService)

	err := rpc.Register(rcvr)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	logger.Log(fmt.Sprintf("Listening tcp network at %d port...", env.Port))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", env.Port))
	if err != nil {
		panic(err)
	}

	logger.Log("Serving grpc...")
	return http.Serve(listener, nil)
}
