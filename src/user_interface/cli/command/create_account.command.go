package command

import (
	"credens/src/application/create"
	"credens/src/domain/account"
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

func NewCreateAccountCommand(commandHandler create.CreateAccountCommandHandler) *cobra.Command {
	type dataParsed struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	validate := validator.New()
	argsParsed := new(dataParsed)

	return &cobra.Command{
		Use:   "create_account json(name, username, password)",
		Short: "Create an account",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires 'json' argument")
			}

			if err := json.Unmarshal([]byte(args[0]), argsParsed); err != nil {
				return err
			}

			return validate.Struct(argsParsed)
		},
		Run: func(cmd *cobra.Command, args []string) {
			commandHandler.Execute(
				*create.NewCreateAccountCommand(
					account.NewAccountId(nil).Value(),
					argsParsed.Name,
					argsParsed.Username,
					argsParsed.Password,
				),
			)
		},
	}
}
