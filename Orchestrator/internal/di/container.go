package di

type Container struct {
	OrchestratorContainer *OrchestratorContainer
	ClientContainer       *ClientContainer
}

var GlobalContainer *Container
