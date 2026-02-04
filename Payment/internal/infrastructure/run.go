package infrastructure

import (
	"go.uber.org/zap"

	"go-link/payment/global"
	"go-link/payment/internal/di"
)

func Run() error {
	LoadConfig()
	SetupLogger()
	di.SetupDependencies()

	grpcServer := NewGRPCServer()
	go func() {
		if err := grpcServer.Run(); err != nil {
			global.LoggerZap.Fatal("gRPC Server failed", zap.Error(err))
		}
	}()

	return nil
}
