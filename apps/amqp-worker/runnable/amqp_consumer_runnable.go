package runnable

import (
	"credens/apps/amqp-worker/config"
	sharedInfraQueue "credens/libs/shared/infrastructure/queue"
	"github.com/defval/inject"
	"log"
)

type AMQPConsumerRunnable struct {
}

func NewAMQPConsumerRunnable() *AMQPConsumerRunnable {
	return &AMQPConsumerRunnable{}
}

func (_ AMQPConsumerRunnable) Run(container *inject.Container, env config.Environment) error {
	var consumer sharedInfraQueue.Consumer
	if err := container.Extract(&consumer); err != nil {
		return err
	}

	data, err := consumer.Consume("default")
	if err != nil {
		return nil
	}

	log.Printf("Consumed: %s\n", string(data))
	return nil
}
