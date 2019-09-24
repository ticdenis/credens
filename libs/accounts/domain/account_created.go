package domain

import (
	bus2 "credens/libs/shared/domain/bus"
	"time"
)

type AccountCreatedData struct {
	Id       string
	Name     string
	Username string
}

type AccountCreated struct {
	message     bus2.Message
	aggregateId string
	occurredOn  int64
	data        AccountCreatedData
}

func NewAccountCreated(id string, name string, username string) *AccountCreated {
	return &AccountCreated{
		*bus2.NewMessage(bus2.EventMessageType),
		id,
		time.Now().UTC().Unix(),
		AccountCreatedData{id, name, username},
	}
}

func (event AccountCreated) Message() bus2.Message {
	return event.message
}

func (event AccountCreated) AggregateId() string {
	return event.aggregateId
}

func (event AccountCreated) EventId() string {
	return event.message.MessageId
}

func (event AccountCreated) EventName() string {
	return "account_created"
}

func (event AccountCreated) EventOccurredOn() int64 {
	return event.occurredOn
}

func (event AccountCreated) Data() interface{} {
	return event.data
}
