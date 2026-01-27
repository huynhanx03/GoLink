package registration

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

// StepCreateUser implements the Identity user creation step.
type StepCreateUser struct {
	State          *State
	IdentityClient ports.IdentityClient
}

func (s *StepCreateUser) Name() string { return "CreateUser" }

func (s *StepCreateUser) Execute(ctx context.Context) error {
	req := dto.CreateUserRequest{
		Username:  s.State.Username,
		Password:  s.State.Password,
		FirstName: s.State.FirstName,
		LastName:  s.State.LastName,
		Gender:    s.State.Gender,
		Birthday:  s.State.Birthday,
	}

	resp, err := s.IdentityClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	s.State.UserID = resp.UserID
	s.State.TenantID = resp.TenantID
	return nil
}

func (s *StepCreateUser) Compensate(ctx context.Context) error {
	if s.State.UserID == 0 {
		return nil
	}
	return s.IdentityClient.DeleteUser(ctx, s.State.UserID)
}
