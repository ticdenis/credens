package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/infrastructure/logging/logrus"
	"credens/src/shared/user_interface/config"
	"credens/src/user_interface/cli/command"
	"github.com/spf13/cobra"
	"os"
)

type Kernel struct {
	env       config.Env
	debug     config.Debug
	container *Container
	rootCmd   *cobra.Command
}

type Container struct {
	Logger   logging.Logger
	Commands []*cobra.Command
}

func NewKernel(environment config.Env, debug config.Debug) *Kernel {
	container := &Container{
		logrus.NewLogger(),
		[]*cobra.Command{},
	}

	container.Commands = append(container.Commands, command.NewHelloCommand(container.Logger))

	rootCmd := new(cobra.Command)
	rootCmd.AddCommand(container.Commands...)

	return &Kernel{
		environment,
		debug,
		container,
		rootCmd,
	}
}

func (kernel *Kernel) Run(args ...string) {
	if len(args) > 0 {
		kernel.rootCmd.SetArgs(args)
	}

	if err := kernel.rootCmd.Execute(); err != nil {
		kernel.container.Logger.Log(err.Error())
		os.Exit(1)
	}
}
