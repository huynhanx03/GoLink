package infrastructure

import (
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"go-link/common/pkg/grpc/interceptors"
	"go-link/notification/global"
	grpcAdapter "go-link/notification/internal/adapters/driver/grpc"
	"go-link/notification/internal/di"
)

// GRPCServer wraps the gRPC server and its configured port.
type GRPCServer struct {
	server *grpc.Server
	port   int
}

// NewGRPCServer creates a gRPC server with auth + error interceptors and
// registers the NotificationService routes from the global DI container.
func NewGRPCServer() *GRPCServer {
	cfg := global.Config
	c := di.GlobalContainer.NotificationContainer

	routes := grpcAdapter.V1Routes(c.Service, c.Repo, c.UnreadCounter)

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.ServerAuthInterceptor(),
			interceptors.ServerErrorInterceptor(),
		),
	)
	routes(srv)

	return &GRPCServer{
		server: srv,
		port:   cfg.Server.GRPCPort,
	}
}

// Run starts the gRPC server and blocks until it stops.
func (s *GRPCServer) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on grpc port %d: %w", s.port, err)
	}

	global.LoggerZap.Info("gRPC server starting", zap.Int("port", s.port))

	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("grpc serve error: %w", err)
	}

	return nil
}

// Stop gracefully shuts down the gRPC server.
func (s *GRPCServer) Stop() {
	global.LoggerZap.Info("Stopping gRPC server...")
	s.server.GracefulStop()
	global.LoggerZap.Info("gRPC server stopped")
}
