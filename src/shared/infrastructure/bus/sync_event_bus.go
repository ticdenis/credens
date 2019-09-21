package bus

import "credens/src/shared/domain/bus"

type SyncEventBus struct {
	eventSubscribers []bus.EventSubscriber
}

func NewSyncEventBus(eventSubscribers []interface{}) bus.EventBus {
	var handlers []bus.EventSubscriber

	for _, handler := range eventSubscribers {
		if _, ok := handler.(bus.EventSubscriber); ok {
			handlers = append(handlers, handler.(bus.EventSubscriber))
		}
	}

	return &SyncEventBus{handlers}
}

func (bus *SyncEventBus) Notify(event bus.Event) {
	for _, subscriber := range bus.eventSubscribers {
		for _, eventName := range subscriber.SubscribedTo() {
			if eventName == event.EventName {
				subscriber.Execute(event)
			}
		}
	}
}

type InMemoryEventPublisher struct {
	events []bus.Event
}

func NewInMemoryEventPublisher() *InMemoryEventPublisher {
	return &InMemoryEventPublisher{[]bus.Event{}}
}

func (publisher InMemoryEventPublisher) Record(domainEvents ...bus.Event) {
	publisher.events = append(publisher.events, domainEvents...)
}

func (publisher InMemoryEventPublisher) Publish(domainEvents ...bus.Event) {
	publisher.events = []bus.Event{}
}
