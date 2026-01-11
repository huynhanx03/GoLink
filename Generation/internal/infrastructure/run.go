package infrastructure

import (
	"go-link/generation/internal/di"
)

func Run() error {
	LoadConfig()
	SetupLogger()
	SetupTimer()
	SetupRedis()
	SetupWideColumn()
	di.SetupDependencies()
	http := NewHTTPServer()

	return http.Run()
}
