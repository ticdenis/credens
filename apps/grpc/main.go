package main

import (
	"credens/apps/grpc/config"
	"credens/apps/grpc/service"
	"credens/libs/shared/infrastructure/logging"
	"fmt"
	"github.com/defval/inject"
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

	logger.Log(fmt.Sprintf("Listening tcp network at %d port...", env.Port))
	server, listener := service.NewGRPCAPIServiceServer(env.Port)

	logger.Log("Serving grpc...")
	return server.Serve(listener)
}
