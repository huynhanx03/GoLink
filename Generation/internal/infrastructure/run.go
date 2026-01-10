package infrastructure

import (
	"go-link/generation/internal/di"
)

func Run() error {
	LoadConfig()
	SetupLogger()
	SetupWideColumn()
	di.SetupDependencies()
	http := NewHTTPServer()

	return http.Run()
}
