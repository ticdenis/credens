package config

import (
	"github.com/defval/inject"

	sharedInfraQueue "credens/libs/shared/infrastructure/queue"
	sharedInfraQueueRabbitMQ "credens/libs/shared/infrastructure/queue/rabbitmq"
)

func BuildContainer(env Environment) (*inject.Container, error) {
	return inject.New(
		inject.Provide(
			sharedInfraQueueRabbitMQ.NewRabbitMQConsumer(
				*sharedInfraQueueRabbitMQ.NewRabbitMQConfig(env.Amqp.Url, "default"),
			),
			inject.As(new(sharedInfraQueue.Consumer)),
		),
	)
}
