package infrastructure

import (
	"fmt"
	"net"

	"go-link/payment/global"
	driverGrpc "go-link/payment/internal/adapters/driver/grpc"
	"go-link/payment/internal/di"
	"go-link/common/pkg/grpc/interceptors"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
	port   int
}

func NewGRPCServer() *GRPCServer {
	cfg := global.Config
	paymentService := di.GlobalContainer.PaymentContainer.Service

	serverRoutes := driverGrpc.V1Routes(paymentService)
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.ServerAuthInterceptor()),
	)

	serverRoutes(srv)

	return &GRPCServer{
		server: srv,
		port:   cfg.Server.GRPCPort,
	}
}

func (s *GRPCServer) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", s.port, err)
	}

	global.LoggerZap.Info("gRPC Server starting", zap.Int("port", s.port))
	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC: %w", err)
	}

	return nil
}

func (s *GRPCServer) Stop() {
	global.LoggerZap.Info("Stopping gRPC Server...")
	s.server.GracefulStop()
	global.LoggerZap.Info("gRPC Server stopped")
}
