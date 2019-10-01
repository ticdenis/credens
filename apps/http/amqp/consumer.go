package amqp

import (
	"credens/apps/http/config"
	infra "credens/libs/shared/infrastructure"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func RunConsumer(env config.Environment) {
	conn, err := amqp.Dial(env.Amqp.Url)
	infra.PanicIfError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	infra.PanicIfError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("custom", true, false, false, false, nil)
	infra.PanicIfError(err, "Could not declare `custom` queue")

	err = amqpChannel.Qos(1, 0, false)
	infra.PanicIfError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	infra.PanicIfError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			customMessage := new(CustomMessage)

			err := json.Unmarshal(d.Body, customMessage)
			infra.PanicIfError(err, "Error decoding JSON")

			log.Printf("Result of %d + %d is : %d",
				customMessage.Numerator, customMessage.Denominator, customMessage.Numerator+customMessage.Denominator,
			)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}
