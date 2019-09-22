package main

func main() {
	config := Bootstrap()

	kernel := NewKernel(config.Env, config.Debug, config.Host, config.Port, config.TimeoutSeconds)

	kernel.Run()
}
