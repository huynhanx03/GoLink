package grpc

import (
	"context"

	identityv1 "go-link/common/gen/go/identity/v1"
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/ports"
)

type IdentityServer struct {
	identityv1.UnimplementedIdentityServiceServer
	userService   ports.UserService
	authService   ports.AuthenticationService
	tenantService ports.TenantService
}

func NewIdentityServer(
	userService ports.UserService,
	authService ports.AuthenticationService,
	tenantService ports.TenantService,
) *IdentityServer {
	return &IdentityServer{
		userService:   userService,
		authService:   authService,
		tenantService: tenantService,
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

func (s *IdentityServer) CreateUser(ctx context.Context, req *identityv1.CreateUserRequest) (*identityv1.CreateUserResponse, error) {
	user, err := s.authService.CreateUser(ctx, &dto.CreateUserRequest{
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    int(req.Gender),
		Birthday:  req.Birthday,
	})
	if err != nil {
		return nil, err
	}

	return &identityv1.CreateUserResponse{
		UserId: int64(user.ID),
	}, nil
}

func (s *IdentityServer) DeleteUser(ctx context.Context, req *identityv1.DeleteUserRequest) (*identityv1.DeleteUserResponse, error) {
	if err := s.userService.Delete(ctx, int(req.UserId)); err != nil {
		return nil, err
	}

	return &identityv1.DeleteUserResponse{
		Success: true,
	}, nil
}

func (s *IdentityServer) UpdateTenantPlan(ctx context.Context, req *identityv1.UpdateTenantPlanRequest) (*identityv1.UpdateTenantPlanResponse, error) {
	planID := int(req.PlanId)
	_, err := s.tenantService.Update(ctx, int(req.TenantId), &dto.UpdateTenantRequest{
		PlanID: &planID,
	})
	if err != nil {
		return nil, err
	}

	return &identityv1.UpdateTenantPlanResponse{
		Success: true,
	}, nil
}
