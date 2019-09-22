package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	pb "credens/src/user_interface/grpc/proto"
	"credens/src/user_interface/grpc/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
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
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", kernel.port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterAPIServiceServer(server, &service.GRPCAPIService{})

	logger.Log("Serving grpc...")
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
