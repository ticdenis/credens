package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/infrastructure/logging/fmt"
	"credens/src/shared/user_interface/config"
	"credens/src/user_interface/cli/command"
	"github.com/spf13/cobra"
	"os"
)

type Kernel struct {
	env       config.Env
	debug     config.Debug
	container *Container
}

type Container struct {
	Logger  logging.Logger
	rootCmd cobra.Command
}

func NewKernel(environment config.Env, debug config.Debug) *Kernel {
	return &Kernel{
		environment,
		debug,
		makeContainer(environment, debug),
	}
}

func (kernel *Kernel) Run(args ...string) {
	if len(args) > 0 {
		kernel.container.rootCmd.SetArgs(args)
	}

	if err := kernel.container.rootCmd.Execute(); err != nil {
		kernel.container.Logger.Log(err.Error())
		os.Exit(1)
	}
}

func makeContainer(environment config.Env, debug config.Debug) *Container {
	container := &Container{
		fmt.NewLogger(),
		cobra.Command{},
	}

	container.rootCmd.AddCommand(
		command.NewHelloCommand(container.Logger),
	)

	return container
}
