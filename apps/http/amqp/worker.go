package amqp

import (
	"credens/apps/http/config"
	infra "credens/libs/shared/infrastructure"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func RunWorker(env config.Environment) {
	conn, err := amqp.Dial(env.Amqp.Url)
	infra.PanicIfError(err, "Can't connect to AMQP: "+env.Amqp.Url)
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	infra.PanicIfError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("custom", true, false, false, false, nil)
	infra.PanicIfError(err, "Could not declare `custom` queue")

	rand.Seed(time.Now().UnixNano())

	customMessage := NewCustomMessage(20, 5)
	body, err := json.Marshal(customMessage)
	infra.PanicIfError(err, "Error encoding JSON")

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	infra.PanicIfError(err, "Error publishing message")

	log.Printf("customMessage: %d+%d", customMessage.Numerator, customMessage.Denominator)
}
