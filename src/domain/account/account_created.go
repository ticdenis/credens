package account

import (
	"credens/src/shared/domain/bus"
)

type AccountCreatedData struct {
	Id       string
	Name     string
	Username string
}

type AccountCreated struct {
	bus.Event
	Data AccountCreatedData
}

func NewAccountCreated(id string, name string, username string) *AccountCreated {
	return &AccountCreated{
		*bus.NewEvent(id, "account_created"),
		AccountCreatedData{id, name, username},
	}
}
