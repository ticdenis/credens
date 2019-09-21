package command

import (
	"credens/src/application/read"
	"credens/src/infrastructure/logging"
	"credens/src/shared/domain/bus"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func NewReadAccountCommand(queryBus bus.QueryBus, logger logging.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "read_account id",
		Short: "Read an account",
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
