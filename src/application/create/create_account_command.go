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

func NewCreateAccountCommand(id string, name string, username string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		*bus.NewCommand("create_account"),
		CreateAccountCommandData{id, name, username, password},
	}
}