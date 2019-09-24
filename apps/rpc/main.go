package main

import (
	"credens/apps/rpc/service"
	"credens/libs/shared/infrastructure/di"
	"credens/libs/shared/infrastructure/logging"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	env, err := LoadEnvironment()
	if err != nil {
		panic(err)
	}

	container := BuildContainer(*env)

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func run(container *di.Container, env Environment) error {
	logger := container.Get(LoggerKey).(logging.Logger)

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
