package account

import (
	"credens/src/shared/domain/bus"
)

type AccountCreated struct {
	bus.Event
	id       string
	name     string
	username string
}

func NewAccountCreated(id string, name string, username string) *AccountCreated {
	return &AccountCreated{
		*bus.NewEvent(id, "account_created"),
		id,
		name,
		username,
	}
}

func (event *AccountCreated) Name() string {
	return event.name
}

func (event *AccountCreated) Username() string {
	return event.username
}
