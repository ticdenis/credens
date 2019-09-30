// make build-run-app name=http

package main

import (
	"credens/apps/http/config"
	"credens/apps/http/runnable"
	infra "credens/libs/shared/infrastructure"
	"github.com/defval/inject"
	"github.com/pkg/errors"
)

func main() {
	env, err := config.LoadEnvironment()
	infra.PanicIfError(err, "Error loading environment!")

	container, err := config.BuildContainer(*env)
	infra.PanicIfError(err, "Error building container!")

	err = run(container, *env)
	infra.PanicIfError(err, "Error running app!")
}

func run(container *inject.Container, env config.Environment) error {
	if err := runnable.NewSQLDatabaseRunnable().Run(container, env); err != nil {
		return errors.Wrap(err, "Error running SQL database!")
	}

	if err := runnable.NewSQLMigrationRunnable().Run(container, env); err != nil {
		return errors.Wrap(err, "Error running SQL migrations!")
	}

	if err := runnable.NewHttpServerRunnable().Run(container, env); err != nil {
		return errors.Wrap(err, "Error running HTTP server!")
	}

	return nil
}
