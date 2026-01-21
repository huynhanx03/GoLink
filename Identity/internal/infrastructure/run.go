package infrastructure

import (
	"go-link/identity/internal/di"
)

// Run starts the Identity service.
func Run() error {
	LoadConfig()
	SetupLogger()
	SetupEnt()
	SetupCache()
	di.SetupDependencies()

	Initialized()

	http := NewHTTPServer()
	return http.Run()
}
