package main

import (
	"credens/apps/http/config"
	"credens/apps/http/server"
	"credens/libs/shared/infrastructure/di"
)

func main() {
	env, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	container := config.BuildContainer(*env)

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func run(container *di.Container, env config.Environment) error {
	return server.NewServer(env, container).Run(env.Port)
}
