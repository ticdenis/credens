// make build-run-app name=http

package main

import (
	"github.com/defval/inject"

	"credens/libs/shared/infrastructure/persistence"

	"credens/apps/http/config"
	"credens/apps/http/server"
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
	err := runSQLDatabase(container)
	if err != nil {
		return err
	}

	return runHTTPServer(env, container)
}

func runSQLDatabase(container *inject.Container) error {
	var sql db.SQLDb
	if err := container.Extract(&sql); err != nil {
		return err
	}
	return sql.Run()
}

func runHTTPServer(env config.Environment, container *inject.Container) error {
	svc, err := server.NewServer(env, container)
	if err != nil {
		return err
	}
	return svc.Run(env.Port)
}
