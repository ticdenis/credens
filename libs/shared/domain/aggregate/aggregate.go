package aggregate

import (
	"credens/libs/shared/domain/bus"
)

type AggregateRoot struct {
	domainEvents []bus.Event
}

func (aggregateRoot *AggregateRoot) PullDomainEvents() []bus.Event {
	events := aggregateRoot.domainEvents

	aggregateRoot.domainEvents = []bus.Event{}

	return events
}

func (aggregateRoot *AggregateRoot) RecordDomainEvent(event bus.Event) {
	aggregateRoot.domainEvents = append(aggregateRoot.domainEvents, event.(bus.Event))
}
