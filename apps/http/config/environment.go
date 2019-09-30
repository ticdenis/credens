package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Environment struct {
	Env   string
	Debug bool
	Port  int
	Sql   struct {
		Driver   string
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
}

func LoadEnvironment() (*Environment, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	environment := new(Environment)

	env, envExists := os.LookupEnv("APP_ENV")
	if !envExists || env == "" {
		env = "development"
	}
	environment.Env = env

	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		return nil, err
	}
	environment.Debug = debug

	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, err
	}
	environment.Port = port

	sqlDriver, sqlDriverExists := os.LookupEnv("DB_DRIVER")
	if !sqlDriverExists || sqlDriver == "" {
		sqlDriver = "mysql"
	}
	environment.Sql.Driver = sqlDriver

	sqlUser, sqlUserExists := os.LookupEnv("MYSQL_USER")
	if !sqlUserExists || sqlUser == "" {
		return nil, errors.New("MYSQL_USER env required!")
	}
	environment.Sql.User = sqlUser

	sqlPassword, sqlPasswordExists := os.LookupEnv("MYSQL_PASSWORD")
	if !sqlPasswordExists || sqlPassword == "" {
		return nil, errors.New("MYSQL_PASSWORD env required!")
	}
	environment.Sql.Password = sqlPassword

	sqlHost, sqlHostExists := os.LookupEnv("MYSQL_HOST")
	if !sqlHostExists || sqlHost == "" {
		return nil, errors.New("MYSQL_HOST env required!")
	}
	environment.Sql.Host = sqlHost

	sqlPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, err
	}
	environment.Sql.Port = sqlPort

	sqlDatabase, sqlDatabaseExists := os.LookupEnv("MYSQL_DATABASE")
	if !sqlDatabaseExists || sqlDatabase == "" {
		return nil, errors.New("MYSQL_DATABASE env required!")
	}
	environment.Sql.Database = sqlDatabase

	return environment, nil
}
