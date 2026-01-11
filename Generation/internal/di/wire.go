package di

func SetupDependencies() *Container {
	container := &Container{
		LinkContainer: InitLinkDependencies(),
	}
	GlobalContainer = container
	return container
}
