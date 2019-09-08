package main

import "credens/src/shared/user_interface"

func main() {
	config := user_interface.Bootstrap()

	kernel := NewKernel(config.Env, config.Debug)

	commandName := "hello"
	args := []string{commandName, "Go"}

	kernel.Run(args...)
}
