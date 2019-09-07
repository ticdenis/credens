package config

import (
	"strconv"
)

type Env string

const (
	LOCAL       Env = "local"
	DEVELOPMENT Env = "development"
	STAGING     Env = "staging"
	PRODUCTION  Env = "production"
)

func ParseEnv(env string) Env {
	return Env(env)
}

type Debug bool

func ParseDebug(str string) Debug {
	debug, _ := strconv.ParseBool(str)

	return Debug(debug)
}
