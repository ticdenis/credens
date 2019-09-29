package command

import (
	"github.com/spf13/cobra"
)

type Command interface {
	Execute() error
}

func NewRootCmd(commands []Command) *cobra.Command {
	rootCmd := new(cobra.Command)

	for _, cmd := range commands {
		if cmd, ok := cmd.(*cobra.Command); ok {
			rootCmd.AddCommand(cmd)
		}
	}

	return rootCmd
}
