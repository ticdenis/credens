// make build-run-app name=http

package main

import (
	"credens/apps/http/config"
	"credens/apps/http/server"
	"database/sql"
	"fmt"
	"github.com/defval/inject"
	"github.com/go-sql-driver/mysql"
)

func main() {
	env, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	db, err := getDB(env)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	container, err := config.BuildContainer(*env, db)
	if err != nil {
		panic(err)
	}

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func getDB(env *config.Environment) (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 env.Sql.User,
		Passwd:               env.Sql.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", env.Sql.Host, env.Sql.Port),
		DBName:               env.Sql.Database,
		AllowNativePasswords: true,
	}
	fmt.Println(cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func run(container *inject.Container, env config.Environment) error {
	svc, err := server.NewServer(env, container)
	if err != nil {
		return err
	}

	return svc.Run(env.Port)
}
