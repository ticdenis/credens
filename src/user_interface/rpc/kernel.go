package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"credens/src/user_interface/rpc/service"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Kernel struct {
	env       config.Env
	debug     config.Debug
	port      int
	container *user_interface.Container
}

func NewKernel(env config.Env, debug config.Debug, port int) *Kernel {
	return &Kernel{
		env,
		debug,
		port,
		NewContainer(env, debug, port),
	}
}

func (kernel *Kernel) Run() {
	logger := kernel.container.Get(LoggerKey).(logging.Logger)

	rcvr := new(service.RPCAPIService)

	err := rpc.Register(rcvr)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	logger.Log(fmt.Sprintf("Listening tcp network at %d port...", kernel.port))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", kernel.port))
	if err != nil {
		panic(err)
	}

	logger.Log("Serving grpc...")
	err = http.Serve(listener, nil)
	if err != nil {
		panic(err)
	}
}
