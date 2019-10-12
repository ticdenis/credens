package domain

import (
	"credens/libs/shared/domain/bus"
)

type (
	AccountCreated struct {
		bus.BaseEvent
	}

	AccountCreatedData struct {
		Id       string
		Name     string
		Username string
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
