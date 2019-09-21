package main

import (
	"credens/src/infrastructure/logging"
	"credens/src/shared/user_interface"
	"credens/src/shared/user_interface/config"
	"github.com/spf13/cobra"
	"os"
)

type Kernel struct {
	env       config.Env
	debug     config.Debug
	container *user_interface.Container
}

func NewKernel(env config.Env, debug config.Debug) *Kernel {
	return &Kernel{
		env,
		debug,
		NewContainer(env, debug),
	}
}

func (kernel *Kernel) Run(args ...string) {
	rootCmd := kernel.container.Get(RootCmdKey).(cobra.Command)

	if len(args) > 0 {
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		kernel.container.Get(LoggerKey).(logging.Logger).Log(err.Error())
		os.Exit(1)
	}
}
