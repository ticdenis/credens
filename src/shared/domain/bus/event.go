package bus

import (
	"credens/src/shared/domain/value_object"
	"time"
)

type Event struct {
	messageId       string
	messageType     string
	aggregateId     string
	eventName       string
	eventOccurredOn int64
}

func NewEvent(aggregateId string, eventName string) *Event {
	return &Event{
		value_object.UUID(nil).Value(),
		"event",
		aggregateId,
		eventName,
		time.Now().UTC().Unix(),
	}
}

func (event *Event) MessageId() string {
	return event.messageId
}

func (event *Event) MessageType() string {
	return "event"
}

func (event *Event) AggregateId() string {
	return event.aggregateId
}

func (event *Event) EventId() string {
	return event.messageId
}

func (event *Event) EventName() string {
	return event.eventName
}

func (event *Event) EventOccurredOn() int64 {
	return event.eventOccurredOn
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
