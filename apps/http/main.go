package main

import (
	"credens/libs/shared/infrastructure/di"
	"credens/libs/shared/infrastructure/logging"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	env, err := LoadEnvironment()
	if err != nil {
		panic(err)
	}

	container := BuildContainer(*env)

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func run(container *di.Container, env Environment) error {
	logger := container.Get(LoggerKey).(logging.Logger)

	address := fmt.Sprintf("%s:%d", env.Host, env.Port)

	server := http.Server{
		Handler:      container.Get(HttpRouterKey).(*mux.Router),
		Addr:         address,
		WriteTimeout: env.TimeoutSeconds * time.Second,
		ReadTimeout:  env.TimeoutSeconds * time.Second,
	}

	logger.Log(fmt.Sprintf("Listening and serving http at %s...", address))

	return server.ListenAndServe()
}
