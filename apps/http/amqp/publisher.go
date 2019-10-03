package amqp

import (
	"credens/apps/http/config"
	queue "credens/libs/shared/infrastructure/queue/data_to_test"
	"credens/libs/shared/infrastructure/queue/rabbitmq"
	"encoding/json"
	"log"
)

func RunPublisher(env config.Environment) {
	publisher := rabbitmq.NewRabbitMQPublisher(*rabbitmq.NewRabbitMQConfig(
		env.Amqp.Url,
		"default",
	))

	publish(publisher, 5, 20)
	publish(publisher, 3, 1)

	log.Println("Nothing to publish, bye!")
}

func publish(publisher *rabbitmq.RabbitMQPublisher, numerator, denominator int) {
	customMessage := queue.NewCustomMessage(numerator, denominator)
	data, _ := json.Marshal(customMessage)
	if err := publisher.Publish(data); err != nil {
		panic(err)
	}
	log.Printf("Published: %+v\n", customMessage)
}
