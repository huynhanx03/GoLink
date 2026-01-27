package interceptors

import (
	"context"

	"google.golang.org/grpc"

	"go-link/common/pkg/grpc/metadata"
)

// ClientAuthInterceptor injects user context from Go context into gRPC metadata.
func ClientAuthInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = metadata.EnsureOutgoingContext(ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// ServerAuthInterceptor extracts user context from gRPC metadata into Go context.
func ServerAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = metadata.ExtractIncomingContext(ctx)
		return handler(ctx, req)
	}
}
