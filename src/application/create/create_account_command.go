package create

import (
	"credens/src/shared/domain/bus"
)

type CreateAccountCommandData struct {
	Id       string
	Name     string
	Username string
	Password string
}

type CreateAccountCommand struct {
	message bus.Message
	data    CreateAccountCommandData
}

func NewCreateAccountCommand(id string, name string, username string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		*bus.NewMessage(bus.CommandMessageType),
		CreateAccountCommandData{id, name, username, password},
	}
}

func (command CreateAccountCommand) Message() bus.Message {
	return command.message
}

func (command CreateAccountCommand) CommandName() string {
	return "create_account"
}

func (command CreateAccountCommand) Data() interface{} {
	return command.data
}
