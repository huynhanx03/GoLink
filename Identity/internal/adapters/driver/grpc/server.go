package grpc

import (
	identityv1 "go-link/common/gen/go/identity/v1"
	"go-link/identity/internal/ports"

	"google.golang.org/grpc"
)

// V1Routes registers the identity service routes
func V1Routes(
	userService ports.UserService,
) func(srv *grpc.Server) {
	return func(srv *grpc.Server) {
		identityv1.RegisterIdentityServiceServer(srv, NewIdentityServer(userService))
	}
}
