package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Kernel struct {
	env            config.Env
	debug          config.Debug
	host           string
	port           int
	timeoutSeconds time.Duration
	container      *user_interface.Container
}

func NewKernel(env config.Env, debug config.Debug, host string, port int, timeoutSeconds time.Duration) *Kernel {
	return &Kernel{
		env,
		debug,
		host,
		port,
		timeoutSeconds,
		NewContainer(env, debug, host, port),
	}
}

func (kernel *Kernel) Run() {
	logger := kernel.container.Get(LoggerKey).(logging.Logger)

	address := fmt.Sprintf("%s:%d", kernel.host, kernel.port)
	timeout := kernel.timeoutSeconds * time.Second

	server := http.Server{
		Handler:      kernel.container.Get(HttpRouterKey).(*mux.Router),
		Addr:         address,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	logger.Log(fmt.Sprintf("Listening and serving http at %s...", address))
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
