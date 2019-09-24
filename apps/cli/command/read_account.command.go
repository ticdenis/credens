package command

import (
	"credens/libs/accounts/application/read"
	"credens/libs/shared/domain/bus"
	"credens/libs/shared/infrastructure/logging"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func NewReadAccountCommand(queryBus bus.QueryBus, logger logging.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "read_account id",
		Short: "Read an accounts",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires 'id' argument")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := queryBus.Ask(*read.NewReadAccountQuery(args[0]))
			if err != nil {
				return err
			}

			logger.Log(fmt.Sprintf("%v", res))

			return nil
		},
	}
}
