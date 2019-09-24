package main

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Environment struct {
	Env   string
	Debug bool
}

func LoadEnvironment() (*Environment, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	env := new(Environment)

	env.Env = os.Getenv("APP_ENV")

	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		return nil, err
	}
	env.Debug = debug

	return env, nil
}
