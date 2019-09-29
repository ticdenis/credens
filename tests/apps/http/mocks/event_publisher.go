package mocks

import (
	"credens/libs/shared/domain/bus"
	"github.com/stretchr/testify/mock"
)

type EventPublisherMock struct {
	mock.Mock
}

func (mock *EventPublisherMock) Record(domainEvents ...bus.Event) {
	mock.Called(domainEvents)
}

func (mock *EventPublisherMock) Publish(domainEvents ...bus.Event) {
	mock.Called(domainEvents)
}
