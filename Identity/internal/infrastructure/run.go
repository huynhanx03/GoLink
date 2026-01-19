package infrastructure

import (
	"go-link/identity/internal/di"
)

// Run starts the Identity service.
func Run() error {
	LoadConfig()
	SetupLogger()
	SetupEnt()
	di.SetupDependencies()

	http := NewHTTPServer()
	return http.Run()
}
