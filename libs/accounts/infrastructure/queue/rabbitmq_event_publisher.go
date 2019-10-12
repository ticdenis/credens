package queue

import (
	"credens/libs/shared/domain/bus"
	"credens/libs/shared/infrastructure/queue"
	"encoding/json"
)

type RabbitMQEventPublisher struct {
	publisher queue.Publisher
}

func NewRabbitMQEventPublisher(publisher queue.Publisher) *RabbitMQEventPublisher {
	return &RabbitMQEventPublisher{publisher}
}

func (rabbitmq *RabbitMQEventPublisher) Publish(domainEvents ...bus.Event) error {
	for _, domainEvent := range domainEvents {
		msg := rabbitmq.toMessageQueue(domainEvent)
		data, _ := json.Marshal(msg)
		if err := rabbitmq.publisher.Publish(data); err != nil {
			return err
		}
	}
	return nil
}

func (_ RabbitMQEventPublisher) toMessageQueue(domainEvent bus.Event) queue.QueueMessage {
	return *queue.NewQueueMessage(
		domainEvent.EventName(),
		domainEvent.Message().MessageId,
		domainEvent.Message().MessageType,
		domainEvent.Data(),
	)
}
