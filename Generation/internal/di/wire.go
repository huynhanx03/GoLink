package di

func SetupDependencies() *Container {
	clientContainer := InitClients()
	linkContainer := InitLinkDependencies(clientContainer)

	container := &Container{
		LinkContainer:   linkContainer,
		ClientContainer: clientContainer,
	}
	GlobalContainer = container
	return container
}
