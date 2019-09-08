package main

func main() {
	config := Bootstrap()

	kernel := NewKernel(config.Env, config.Debug)

	commandName := "hello"
	args := []string{commandName, "Go"}

	kernel.Run(args...)
}
