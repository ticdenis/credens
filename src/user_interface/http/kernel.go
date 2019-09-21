package main

import (
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"net/http"
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

func (kernel *Kernel) Run() {
	server := kernel.container.Get(HttpServerKey).(http.Server)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
