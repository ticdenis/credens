package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	config RabbitMQConfig
}

func NewRabbitMQPublisher(config RabbitMQConfig) *RabbitMQPublisher {
	return &RabbitMQPublisher{config: config}
}

func (publisher *RabbitMQPublisher) Publish(message []byte) error {
	connection, err := amqp.Dial(publisher.config.Url)
	if err != nil {
		return errors.Wrap(err, "Can't connect to AMQP")
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return errors.Wrap(err, "Can't create a AMQP Channel")
	}
	defer channel.Close()

	queueDeclared, err := channel.QueueDeclare(
		publisher.config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "Could not declare `%s` queue", publisher.config.QueueName)
	}

	return channel.Publish(
		"",
		queueDeclared.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         message,
		},
	)
}
