package grpc

import (
	"context"

	identityv1 "go-link/common/gen/go/identity/v1"
	"go-link/common/pkg/grpc"
	"go-link/common/pkg/settings"
	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

type identityClientAdapter struct {
	client identityv1.IdentityServiceClient
}

func NewIdentityClient(cfg settings.GRPCService) (ports.IdentityClient, error) {
	conn, err := grpc.NewClientConn(cfg)
	if err != nil {
		return nil, err
	}
	return &identityClientAdapter{
		client: identityv1.NewIdentityServiceClient(conn),
	}, nil
}

func (a *identityClientAdapter) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	resp, err := a.client.CreateUser(ctx, &identityv1.CreateUserRequest{
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    int32(req.Gender),
		Birthday:  req.Birthday,
	})
	if err != nil {
		return dto.CreateUserResponse{}, err
	}
	return dto.CreateUserResponse{
		UserID:   resp.UserId,
		TenantID: resp.TenantId,
	}, nil
}

func (a *identityClientAdapter) DeleteUser(ctx context.Context, userID int64) error {
	_, err := a.client.DeleteUser(ctx, &identityv1.DeleteUserRequest{
		UserId: userID,
	})
	return err
}
