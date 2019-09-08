package create

import "credens/src/shared/domain/bus"

type Data struct {
	Id       string
	Name     string
	UserName string
	Password string
}

type CreateAccountCommand struct {
	bus.Command
	Data Data
}

func NewCreateAccountCommand(id string, name string, userName string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		*bus.NewCommand("create_account"),
		Data{id, name, userName, password},
	}
}
