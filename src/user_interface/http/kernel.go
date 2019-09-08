package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/infrastructure/logging/logrus"
	"credens/src/shared/user_interface/config"
)

type Kernel struct {
	env       config.Env
	debug     config.Debug
	container *Container
}

type Container struct {
	Logger logging.Logger
}

func NewKernel(environment config.Env, debug config.Debug) *Kernel {
	container := &Container{
		logrus.NewLogger(),
	}

	return &Kernel{
		environment,
		debug,
		container,
	}
}

func (kernel *Kernel) Run(req interface{}) {
	if nil != req {
		kernel.container.Logger.Log("HTTP should execute this req: " + req.(string))
	} else {
		kernel.container.Logger.Log("Hello from HTTP!")
	}
}
