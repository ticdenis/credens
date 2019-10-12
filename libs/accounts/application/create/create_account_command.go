package create

import (
	"credens/libs/shared/domain/bus"
)

var commandName = "create_account"

type (
	CreateAccountCommand struct {
		bus.BaseCommand
	}

	CreateAccountCommandData struct {
		Id       string
		Name     string
		Username string
		Password string
	}
)

func NewCreateAccountCommand(data CreateAccountCommandData) *CreateAccountCommand {
	return &CreateAccountCommand{*bus.NewCommand(commandName, data)}
}
