package main

import (
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
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

type API int

type Item struct {
	Text string
}

func (api *API) SayHello(to string, reply *Item) error {
	(*reply).Text = fmt.Sprintf("Hello %s from RPC app!", to)

	return nil
}

func (kernel *Kernel) Run() {
	port := 4040

	var api = new(API)

	err := rpc.Register(api)
	if err != nil {
		log.Fatal("register error", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("listener error", err)
	}

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("serve error", err)
	}

	log.Printf("Serving rpc on port %d", port)
}
