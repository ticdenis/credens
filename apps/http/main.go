// make build-run-app name=http

package main

import (
	"credens/apps/http/config"
	"credens/apps/http/server"
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
	svc, err := server.NewServer(env, container)
	if err != nil {
		return err
	}

	return svc.Run(env.Port)
}
