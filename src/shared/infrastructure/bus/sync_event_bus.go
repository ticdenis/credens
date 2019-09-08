package bus

import "credens/src/shared/domain/bus"

type SyncEventBus struct {
	eventSubscribers []bus.EventSubscriber
}

func NewSyncEventBus(eventSubscribers []bus.EventSubscriber) bus.EventBus {
	return &SyncEventBus{eventSubscribers}
}

func (bus *SyncEventBus) Notify(event bus.Event) {
	for _, subscriber := range bus.eventSubscribers {
		for _, eventName := range subscriber.SubscribedTo() {
			if eventName == event.EventName() {
				subscriber.Execute(event)
			}
		}
	}
}
