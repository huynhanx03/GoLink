package infrastructure

import (
	"context"
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

	di.GlobalContainer.LinkContainer.CodePool.Start(context.Background())

	return http.Run()
}
