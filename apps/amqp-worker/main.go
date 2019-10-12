// make build-run-app name=amqp-worker

package main

import (
	"credens/apps/amqp-worker/config"
	"credens/apps/amqp-worker/runnable"
	infra "credens/libs/shared/infrastructure"
	"github.com/defval/inject"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.Println("Starting amqp-worker app...")

	env, err := config.LoadEnvironment()
	infra.PanicIfError(err, "Error loading environment!")

	container, err := config.BuildContainer(*env)
	infra.PanicIfError(err, "Error building container!")

	log.Println("Consuming...")
	for {
		err = run(container, *env)
		infra.PanicIfError(err, "Error running app!")
	}
}

func run(container *inject.Container, env config.Environment) error {
	if err := runnable.NewAMQPConsumerRunnable().Run(container, env); err != nil {
		return errors.Wrap(err, "Error running AMQP consumer!")
	}
	return nil
}
