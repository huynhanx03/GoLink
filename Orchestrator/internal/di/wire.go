package di

import (
	"go-link/orchestrator/global"
)

func SetupDependencies() *Container {
	clientContainer := InitClients(global.Config)
	orchestratorContainer := InitOrchestratorDependencies(clientContainer)

	container := &Container{
		OrchestratorContainer: orchestratorContainer,
		ClientContainer:       clientContainer,
	}
	GlobalContainer = container
	return container
}
