package main

import "credens/src/shared/user_interface"

func main() {
	config := user_interface.Bootstrap()

	kernel := NewKernel(config.Env, config.Debug)

	// commandName := "create_account"
	// args := []string{commandName, "{\"name\": \"foo\", \"username\": \"alias\", \"password\": \"pwd\"}"}
	var args []string

	kernel.Run(args...)
}
