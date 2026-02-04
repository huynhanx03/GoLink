package infrastructure

import (
	"go-link/orchestrator/internal/di"
)

// Run starts the Orchestrator service.
func Run() error {
	LoadConfig()
	SetupLogger()
	di.SetupDependencies()

	http := NewHTTPServer()
	return http.Run()
}
