package main

func main() {
	config := Bootstrap()

	kernel := NewKernel(config.Env, config.Debug, config.Port)

	kernel.Run()
}
