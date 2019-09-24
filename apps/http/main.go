package main

import (
	"credens/libs/shared/infrastructure/di"
	"fmt"
	"github.com/gin-gonic/gin"
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
	server := container.Get(HttpServerKey).(gin.Engine)

	return server.Run(fmt.Sprintf(":%d", env.Port))
}
