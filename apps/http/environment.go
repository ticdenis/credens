package main

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Environment struct {
	Env   string
	Debug bool
	Port  int
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
	environment.Port = port
	if err != nil {
		return nil, err
	}

	return environment, nil
}
