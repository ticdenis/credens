package main

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Environment struct {
	Env            string
	Debug          bool
	Host           string
	Port           int
	TimeoutSeconds time.Duration
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

	env.Host = os.Getenv("HTTP_HOST")

	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	env.Port = port
	if err != nil {
		return nil, err
	}

	timeoutSeconds, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT_SECONDS"))
	if err != nil {
		return nil, err
	}
	env.TimeoutSeconds = time.Duration(timeoutSeconds)

	return env, nil
}
