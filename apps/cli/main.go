// make build-run-app name=cli args="help"

package main

import (
	"credens/apps/cli/command"
	"credens/apps/cli/config"
	"github.com/defval/inject"
)

func main() {
	env, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	container, err := config.BuildContainer(*env)
	if err != nil {
		panic(err)
	}

	err = run(container, *env)
	if err != nil {
		panic(err)
	}
}

func run(container *inject.Container, env config.Environment) error {
	var commands []command.Command
	if err := container.Extract(&commands); err != nil {
		return err
	}

	return command.NewRootCmd(commands).Execute()
}
