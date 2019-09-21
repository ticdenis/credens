package command

import (
	"credens/src/application/create"
	"credens/src/domain/account"
	"credens/src/shared/domain/bus"
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

func NewCreateAccountCommand(commandBus bus.CommandBus) *cobra.Command {
	var argsParsed struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

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

			return validator.New().Struct(argsParsed)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return commandBus.Dispatch(
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
