package bus

import "time"

type (
	Event interface {
		Message() Message
		AggregateId() string
		EventId() string
		EventName() string
		EventOccurredOn() int64
		Data() interface{}
	}

	BaseEvent struct {
		eventName   string
		message     Message
		aggregateId string
		occurredOn  int64
		data        interface{}
	}

	EventBus interface {
		Notify(event Event)
	}

	EventPublisher interface {
		Record(domainEvents ...Event)
		Publish(domainEvents ...Event)
	}

	EventSubscriber interface {
		SubscribedTo() []string
		Execute(event Event)
	}
)

func NewEvent(eventName string, aggregateId string, data interface{}) *BaseEvent {
	return &BaseEvent{
		message:     *NewMessage("event"),
		aggregateId: aggregateId,
		occurredOn:  time.Now().UTC().Unix(),
		data:        data,
	}
}

func (event BaseEvent) Message() Message {
	return event.message
}

func (event BaseEvent) AggregateId() string {
	return event.aggregateId
}

func (event BaseEvent) EventId() string {
	return event.message.MessageId
}

func (event BaseEvent) EventName() string {
	return event.eventName
}

func (event BaseEvent) EventOccurredOn() int64 {
	return event.occurredOn
}

func (event BaseEvent) Data() interface{} {
	return event.data
}
