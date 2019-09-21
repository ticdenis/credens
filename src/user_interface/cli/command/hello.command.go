package command

import (
	"credens/src/infrastructure/logging"
	"errors"
	"github.com/spf13/cobra"
)

func NewHelloCommand(logger logging.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "hello [who]",
		Short: "Salutation for X",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires 'who' argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			logger.Log("Hello " + args[0] + "!\n")
		},
	}
}
