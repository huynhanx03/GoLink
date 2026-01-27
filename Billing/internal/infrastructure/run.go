package infrastructure

import (
	"fmt"

	"go-link/billing/internal/di"
)

// Run starts the Billing service.
func Run() error {
	LoadConfig()
	SetupLogger()
	SetupEnt()
	SetupRedis()
	SetupCache()
	SetupKeys()
	di.SetupDependencies()

	Initialized()

	// Start gRPC Server
	grpcServer := NewGRPCServer()
	go func() {
		if err := grpcServer.Run(); err != nil {
			panic(fmt.Sprintf("Failed to run gRPC server: %v", err))
		}
	}()

	http := NewHTTPServer()
	return http.Run()
}
