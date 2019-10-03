// make build-run-app name=amqp-worker

package main

import (
	"credens/apps/amqp-worker/config"
	infra "credens/libs/shared/infrastructure"
	queue "credens/libs/shared/infrastructure/queue/data_to_test"
	"credens/libs/shared/infrastructure/queue/rabbitmq"
	"encoding/json"
	"github.com/defval/inject"
	"log"
)

func main() {
	env, err := config.LoadEnvironment()
	infra.PanicIfError(err, "Error loading environment!")

	container, err := config.BuildContainer(*env)
	infra.PanicIfError(err, "Error building container!")

	err = run(container, *env)
	infra.PanicIfError(err, "Error running app!")
}

func run(container *inject.Container, env config.Environment) error {
	consumer := rabbitmq.NewRabbitMQConsumer(*rabbitmq.NewRabbitMQConfig(
		env.Amqp.Url,
		"default",
	))

	for {
		go consume(consumer)
	}
}

func consume(consumer *rabbitmq.RabbitMQConsumer) {
	data, err := consumer.Consume("default")
	if err != nil {
		panic(err)
	}
	customMessage := new(queue.CustomMessage)
	_ = json.Unmarshal(data, &customMessage)
	log.Printf("Consumed: %+v\n", customMessage)
}
