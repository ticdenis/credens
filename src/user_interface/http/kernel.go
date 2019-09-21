package main

import (
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
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
	host := "127.0.0.1"
	port := 8000

	timeout := 15 * time.Second

	server := http.Server{
		Handler:      kernel.container.Get(HttpRouterKey).(*mux.Router),
		Addr:         fmt.Sprintf("%s:%d", host, port),
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error listening and serve server", err)
	}
	log.Printf("Serving http on port %d", port)
}
