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
		Driver  string
		Migrate bool
		Url     string
	}
	Amqp struct {
		Driver string
		Url    string
	}
}

func LoadEnvironment() (*Environment, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	environment := new(Environment)
	loaders := []func() error{
		environment.loadAPPEnvs,
		environment.loadHTTPEnvs,
		environment.loadSQLEnvs,
		environment.loadAMQPEnvs,
	}

	for _, fn := range loaders {
		if err := fn(); err != nil {
			return nil, err
		}
	}

	return environment, nil
}

func (environment *Environment) loadAPPEnvs() error {
	env, envExists := os.LookupEnv("APP_ENV")
	if !envExists || env == "" {
		env = "development"
	}
	environment.Env = env

	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		debug = false
	}
	environment.Debug = debug

	return nil
}

func (environment *Environment) loadHTTPEnvs() error {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return err
	}
	environment.Port = port

	return nil
}

func (environment *Environment) loadSQLEnvs() error {
	sqlDriver, sqlDriverExists := os.LookupEnv("SQL_DRIVER")
	if !sqlDriverExists || sqlDriver == "" {
		return errors.New("SQL_DRIVER env required!")
	}
	environment.Sql.Driver = sqlDriver

	sqlMigrate, err := strconv.ParseBool(os.Getenv("SQL_MIGRATE"))
	if err != nil {
		environment.Sql.Migrate = false
	}
	environment.Sql.Migrate = sqlMigrate

	sqlUrl, sqlUrlExists := os.LookupEnv("SQL_URL")
	if !sqlUrlExists || sqlUrl == "" {
		return errors.New("SQL_URL env required!")
	}
	environment.Sql.Url = sqlUrl

	return nil
}

func (environment *Environment) loadAMQPEnvs() error {
	amqpDriver, amqpDriverExists := os.LookupEnv("AMQP_DRIVER")
	if !amqpDriverExists || amqpDriver == "" {
		return errors.New("AMQP_DRIVER env required!")
	}
	environment.Amqp.Driver = amqpDriver

	amqpUrl, amqpUrlExists := os.LookupEnv("AMQP_URL")
	if !amqpUrlExists || amqpUrl == "" {
		return errors.New("AMQP_URL env required!")
	}
	environment.Amqp.Url = amqpUrl

	return nil
}
