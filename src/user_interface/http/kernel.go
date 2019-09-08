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
	return &Kernel{
		environment,
		debug,
		makeContainer(environment, debug),
	}
}

func (kernel *Kernel) Run(req interface{}) {
	if nil != req {
		kernel.container.Logger.Log("HTTP should execute this req: " + req.(string))
	} else {
		kernel.container.Logger.Log("Hello from HTTP!")
	}
}

func makeContainer(environment config.Env, debug config.Debug) *Container {
	return &Container{
		logrus.NewLogger(),
	}
}