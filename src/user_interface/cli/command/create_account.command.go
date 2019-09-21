package command

import (
	"credens/src/application/create"
	"credens/src/domain/account"
	"errors"
	"github.com/spf13/cobra"
)

func NewCreateAccountCommand(commandHandler create.CreateAccountCommandHandler) *cobra.Command {
	return &cobra.Command{
		Use:   "create_account [name] [username] [password]",
		Short: "Create an account",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires 'name' argument")
			} else if len(args) < 2 {
				return errors.New("requires 'username' argument")
			} else if len(args) < 3 {
				return errors.New("requires 'password' argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			commandHandler.Execute(
				*create.NewCreateAccountCommand(
					account.NewAccountId(nil).Value(),
					args[0],
					args[1],
					args[2],
				),
			)
		},
	}
}
