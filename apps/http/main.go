// make build-run-app name=http

package main

import (
	"credens/apps/http/config"
	"credens/apps/http/migration/sql_migration"
	"credens/apps/http/server"
	"credens/libs/shared/infrastructure/persistence"
	"github.com/defval/inject"
	"github.com/pkg/errors"
)

func main() {
	env, err := config.LoadEnvironment()
	panicIfError(err, "Error loading environment!")

	container, err := config.BuildContainer(*env)
	panicIfError(err, "Error building container!")

	err = run(container, *env)
	panicIfError(err, "Error running app!")
}

func panicIfError(err error, msg string) {
	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}

func run(container *inject.Container, env config.Environment) error {
	if err := runSQLDatabase(container, env); err != nil {
		return errors.Wrap(err, "Error running SQL database!")
	}

	if err := runSQLMigration(container, env); err != nil {
		return errors.Wrap(err, "Error running SQL migrations!")
	}

	if err := runHTTPServer(container, env); err != nil {
		return errors.Wrap(err, "Error running HTTP server!")
	}

	return nil
}

func runSQLDatabase(container *inject.Container, env config.Environment) error {
	var sqlDB db.SQLDb
	if err := container.Extract(&sqlDB); err != nil {
		return err
	}
	return sqlDB.Run();
}

func runSQLMigration(container *inject.Container, env config.Environment) error {
	if env.Sql.Migrate {
		var sqlMigrator sql_migration.SQLMigrator
		if err := container.Extract(&sqlMigrator); err != nil {
			return err
		}
		return sqlMigrator.Run()
	}
	return nil
}

func runHTTPServer(container *inject.Container, env config.Environment) error {
	svc, err := server.NewServer(env, container)
	if err != nil {
		return err
	}
	return svc.Run(env.Port)
}
