package infrastructure

import (
	"context"

	"go.uber.org/zap"

	"go-link/redirection/global"
	"go-link/redirection/internal/di"
)

func Run() error {
	LoadConfig()
	SetupLogger()
	SetupRedis()
	SetupWideColumn()
	di.SetupDependencies()
	http := NewHTTPServer()

	ctx := context.Background()

	if err := di.GlobalContainer.LinkContainer.Consumer.Start(ctx); err != nil {
		global.Logger.Error("Link CDC Consumer failed", zap.Error(err))
	}

	return http.Run()
}
