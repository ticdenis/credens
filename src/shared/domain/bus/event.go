package bus

import (
	"time"
)

type Event struct {
	Message
	AggregateId     string
	EventId         string
	EventName       string
	EventOccurredOn int64
}

var eventMessageType = "event"

func NewEvent(aggregateId string, eventName string) *Event {
	message := *NewMessage(eventMessageType)

	return &Event{
		message,
		aggregateId,
		message.MessageId,
		eventName,
		time.Now().UTC().Unix(),
	}
}

type EventBus interface {
	Notify(event Event)
}

type EventPublisher interface {
	Record(domainEvents ...Event)
	Publish(domainEvents ...Event)
}

type EventSubscriber interface {
	SubscribedTo() []string
	Execute(event Event)
}
