package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	config RabbitMQConfig
}

func NewRabbitMQConsumer(config RabbitMQConfig) *RabbitMQConsumer {
	return &RabbitMQConsumer{config: config}
}

func (consumer *RabbitMQConsumer) Consume(key string) ([]byte, error) {
	connection, err := amqp.Dial(consumer.config.Url)
	if err != nil {
		return nil, errors.Wrap(err, "Can't connect to AMQP")
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Can't create a AMQP Channel")
	}
	defer channel.Close()

	queueDeclared, err := channel.QueueDeclare(
		consumer.config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not declare `%s` queue", consumer.config.QueueName)
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		return nil, errors.Wrap(err, "Could not configure QoS")
	}

	messageChannel, err := channel.Consume(
		queueDeclared.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Could not register consumer")
	}

	var delivery = <-messageChannel
	body := delivery.Body
	if err := delivery.Ack(false); err != nil {
		return nil, err
	}
	return body, nil
}