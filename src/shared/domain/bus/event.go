package bus

type Event interface {
	Message() Message
	AggregateId() string
	EventId() string
	EventName() string
	EventOccurredOn() int64
	Data() interface{}
}

var EventMessageType = "event"

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
