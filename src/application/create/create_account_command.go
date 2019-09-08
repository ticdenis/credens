package create

import "credens/src/shared/domain/bus"

type CreateAccountCommandData struct {
	Id       string
	Name     string
	UserName string
	Password string
}

type CreateAccountCommand struct {
	bus.Command
	Data CreateAccountCommandData
}

func NewCreateAccountCommand(id string, name string, userName string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		*bus.NewCommand("create_account"),
		CreateAccountCommandData{id, name, userName, password},
	}
}
