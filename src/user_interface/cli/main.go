package main

func main() {
	config := Bootstrap()

	kernel := NewKernel(config.Env, config.Debug)

	kernel.Run(handler(&kernel.Container))
}

func handler(container *Container) func() {
	return func() {
		container.Logger.Log("Hello world!")
	}
}
