// make build-run-app name=http

package main

import (
	"credens/apps/http/config"
	"credens/apps/http/migration/sql_migration"
	"credens/apps/http/server"
	"credens/libs/shared/infrastructure/persistence"
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
	err := runSQLDatabase(container, env)
	if err != nil {
		return err
	}

	return runHTTPServer(container, env)
}

func runSQLDatabase(container *inject.Container, env config.Environment) error {
	var sqlDB db.SQLDb
	if err := container.Extract(&sqlDB); err != nil {
		return err
	}

	if err := sqlDB.Run(); err != nil {
		return err
	}

	if env.Sql.Migrate {
		return sql_migration.Migrate(sqlDB.DB())
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
