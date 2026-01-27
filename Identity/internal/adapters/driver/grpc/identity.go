package grpc

import (
	"context"

	identityv1 "go-link/common/gen/go/identity/v1"
	"go-link/identity/internal/ports"
)

type IdentityServer struct {
	identityv1.UnimplementedIdentityServiceServer
	userService ports.UserService
}

func NewIdentityServer(userService ports.UserService) *IdentityServer {
	return &IdentityServer{
		userService: userService,
	}
}

func (s *IdentityServer) GetUserRole(ctx context.Context, req *identityv1.GetUserRoleRequest) (*identityv1.GetUserRoleResponse, error) {
	role, err := s.userService.GetRole(ctx, int(req.UserId), int(req.TenantId))
	if err != nil {
		return nil, err
	}

	return &identityv1.GetUserRoleResponse{
		Role: &identityv1.Role{
			Id:    int64(role.ID),
			Name:  role.Name,
			Level: int32(role.Level),
		},
	}, nil
}
