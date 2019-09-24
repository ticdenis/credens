package create

import (
	bus2 "credens/libs/shared/domain/bus"
)

var commandName = "create_account"

type CreateAccountCommandData struct {
	Id       string
	Name     string
	Username string
	Password string
}

type CreateAccountCommand struct {
	message bus2.Message
	data    CreateAccountCommandData
}

func NewCreateAccountCommand(id string, name string, username string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		*bus2.NewMessage(bus2.CommandMessageType),
		CreateAccountCommandData{id, name, username, password},
	}
}

func (command CreateAccountCommand) Message() bus2.Message {
	return command.message
}

func (command CreateAccountCommand) CommandName() string {
	return commandName
}

func (command CreateAccountCommand) Data() interface{} {
	return command.data
}
