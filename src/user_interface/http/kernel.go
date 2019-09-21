package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
)

type Kernel struct {
	env       config.Env
	debug     config.Debug
	container *user_interface.Container
}

func NewKernel(env config.Env, debug config.Debug) *Kernel {
	return &Kernel{
		env,
		debug,
		NewContainer(env, debug),
	}
}

func (kernel *Kernel) Run(req interface{}) {
	logger := kernel.container.Get(LoggerKey).(logging.Logger)

	if nil != req {
		logger.Log("HTTP should execute this req: " + req.(string))
	} else {
		logger.Log("Hello from HTTP!")
	}
}
