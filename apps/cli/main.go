package main

import (
	"credens/libs/shared/infrastructure/di"
	"github.com/spf13/cobra"
)

func main() {
	env, err := LoadEnvironment()
	if err != nil {
		panic(err)
	}

	container := BuildContainer(*env)

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func run(container *di.Container, env Environment) error {
	rootCmd := container.Get(RootCmdKey).(cobra.Command)

	return rootCmd.Execute()
}
