package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go-link/common/pkg/grpc/interceptors"
	"go-link/common/pkg/settings"
)

// ClientOptions defines the configuration for gRPC client connection
type ClientOptions struct {
	Target       string
	Timeout      time.Duration
	Interceptors []grpc.UnaryClientInterceptor
}

// NewClientConn creates a new gRPC client connection with default secure settings and interceptors
func NewClientConn(service settings.GRPCService, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	// Default options
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUnaryInterceptor(interceptors.ClientAuthInterceptor()),
	}

	// Append user options
	defaultOpts = append(defaultOpts, opts...)

	target := fmt.Sprintf("dns:///%s:%d", service.Host, service.Port)

	// Create connection
	conn, err := grpc.Dial(target, defaultOpts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// NewClientConnWithContext creates a connection with context timeout
func NewClientConnWithContext(ctx context.Context, service settings.GRPCService, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := NewClientConn(service, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
