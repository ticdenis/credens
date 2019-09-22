package main

import (
	"credens/src/shared/user_interface"
	"os"
	"strconv"
	"time"
)

type Config struct {
	user_interface.Config
	Host           string
	Port           int
	TimeoutSeconds time.Duration
}

func Bootstrap() Config {
	config := user_interface.Bootstrap()

	host := os.Getenv("HTTP_HOST")
	if host == "" {
		panic("HTTP_HOST can't be empty!")
	}

	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}

	timeoutSeconds, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT_SECONDS"))
	if err != nil {
		panic(err)
	}

	return Config{
		config,
		host,
		port,
		time.Duration(timeoutSeconds),
	}
}
