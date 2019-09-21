package create

import "credens/src/shared/domain/bus"

type CreateAccountCommandData struct {
	Id       string
	Name     string
	Username string
	Password string
}

type CreateAccountCommand struct {
	bus.Command
	Data CreateAccountCommandData
}

var createAccountCommandName = "create_account"

func NewCreateAccountCommand(id string, name string, username string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		*bus.NewCommand(createAccountCommandName),
		CreateAccountCommandData{id, name, username, password},
	}
}
