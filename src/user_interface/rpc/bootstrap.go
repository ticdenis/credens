package main

import (
	"credens/src/shared/user_interface"
	"os"
	"strconv"
)

type Config struct {
	user_interface.Config
	Port int
}

func Bootstrap() Config {
	config := user_interface.Bootstrap()

	port, err := strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		panic(err)
	}

	return Config{
		config,
		port,
	}
}
