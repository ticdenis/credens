package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/infrastructure/logging/logrus"
	"credens/src/shared/user_interface/config"
)

type Kernel struct {
	Env       config.Env
	Debug     config.Debug
	Container Container
}

type Container struct {
	Logger logging.Logger
}

func NewKernel(environment config.Env, debug config.Debug) *Kernel {
	container := Container{
		logrus.NewLogger(),
	}

	return &Kernel{
		environment,
		debug,
		container,
	}
}

func (kernel *Kernel) Run(handler func()) {
	handler()
}