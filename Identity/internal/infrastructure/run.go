package infrastructure

import (
	"go-link/identity/global"
	"go-link/identity/internal/di"

	"go.uber.org/zap"
)

// Run starts the Identity service.
func Run() error {
	LoadConfig()
	SetupLogger()
	SetupEnt()
	SetupCache()
	SetupKeys()
	di.SetupDependencies()

	Initialized()

	grpcServer := NewGRPCServer()
	go func() {
		if err := grpcServer.Run(); err != nil {
			global.LoggerZap.Fatal("gRPC Server failed", zap.Error(err))
		}
	}()

	http := NewHTTPServer()
	return http.Run()
}
