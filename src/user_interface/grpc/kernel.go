package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"credens/src/user_interface/grpc/service"
	"fmt"
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

	logger.Log(fmt.Sprintf("Listening tcp network at %d port...", kernel.port))
	server, listener := service.NewGRPCAPIServiceServer(kernel.port)

	logger.Log("Serving grpc...")
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
