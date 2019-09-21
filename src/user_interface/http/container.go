package main

import (
	"credens/src/infrastructure/logging/logrus"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
)

const (
	LoggerKey string = "credens/src/infrastructure/logging/Logger"
)

func NewContainer(env config.Env, debug config.Debug) *user_interface.Container {
	ctx := user_interface.NewContainer()

	ctx.Set(LoggerKey, func(_ *user_interface.Container) interface{} {
		return logrus.NewLogger()
	})

	return ctx
}
