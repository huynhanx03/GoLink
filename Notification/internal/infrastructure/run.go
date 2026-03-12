package infrastructure

import (
	"context"
	"time"

	"go.uber.org/zap"

	"os"
	"os/signal"
	"syscall"

	"go-link/notification/global"
	"go-link/notification/internal/di"
)

// Run starts the Notification service: config, logger, storage, DI wiring,
// Kafka consumer, gRPC server, and HTTP server.
func Run() error {
	LoadConfig()
	SetupLogger()
	SetupMongoDB()
	SetupRedis()
	SetupKeys()
	di.SetupDependencies(context.Background())

	// Start consumers
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer := di.GlobalContainer.NotificationContainer.Consumer

	go func() {
		if err := consumer.Start(ctx); err != nil {
			global.LoggerZap.Error("Notification consumer stopped", zap.Error(err))
		}
	}()

	go func() {
		if err := consumer.StartRetryConsumer(ctx); err != nil {
			global.LoggerZap.Error("Notification retry consumer stopped", zap.Error(err))
		}
	}()

	// Start digest worker
	digestWorker := di.GlobalContainer.NotificationContainer.DigestWorker
	go func() {
		if err := digestWorker.Start(ctx); err != nil {
			global.LoggerZap.Error("Digest worker stopped", zap.Error(err))
		}
	}()

	// Start gRPC
	grpcServer := NewGRPCServer()
	go func() {
		if err := grpcServer.Run(); err != nil {
			global.LoggerZap.Error("gRPC server stopped", zap.Error(err))
		}
	}()

	// Start HTTP
	httpServerWrapper := NewHTTPServer()
	httpSrv, err := httpServerWrapper.Start()
	if err != nil {
		global.LoggerZap.Fatal("Failed to start HTTP server", zap.Error(err))
	}

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.LoggerZap.Info("Shutting down server...")

	// Stop consumers
	cancel()

	// Graceful consumer stop
	if err := consumer.Stop(); err != nil {
		global.LoggerZap.Error("Failed to stop consumer", zap.Error(err))
	}
	digestWorker.Stop()

	// Stop servers
	grpcServer.Stop()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		global.LoggerZap.Error("HTTP server forced to shutdown", zap.Error(err))
	}

	// Close resources
	if err := di.GlobalContainer.NotificationContainer.Producer.Close(); err != nil {
		global.LoggerZap.Error("Failed to close producer", zap.Error(err))
	}

	global.LoggerZap.Info("Server exited gracefully")
	return nil
}
