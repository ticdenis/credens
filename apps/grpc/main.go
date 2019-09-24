package main

import (
	"credens/apps/grpc/service"
	"credens/libs/shared/infrastructure/di"
	"credens/libs/shared/infrastructure/logging"
	"fmt"
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

	logger.Log(fmt.Sprintf("Listening tcp network at %d port...", env.Port))
	server, listener := service.NewGRPCAPIServiceServer(env.Port)

	logger.Log("Serving grpc...")
	return server.Serve(listener)
}
