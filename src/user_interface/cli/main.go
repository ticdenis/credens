package main

import "credens/src/shared/user_interface"

func main() {
	config := user_interface.Bootstrap()

	kernel := NewKernel(config.Env, config.Debug)

	kernel.Run()

	// commandName := "read_account"
	// args := []string{commandName, "F7605911-80F5-4F12-8F41-3BD68D77FDF8"}
	//kernel.Run(args...)
}
