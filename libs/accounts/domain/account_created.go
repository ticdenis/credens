package domain

import (
	"credens/libs/shared/domain/bus"
)

type (
	AccountCreatedData struct {
		Id       string
		Name     string
		Username string
	}

	AccountCreated struct {
		bus.BaseEvent
	}
)

func NewAccountCreated(id string, name string, username string) *AccountCreated {
	return &AccountCreated{
		*bus.NewEvent(
			"account_created",
			id,
			AccountCreatedData{id, name, username},
		),
	}
}
