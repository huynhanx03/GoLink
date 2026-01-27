package registration

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/core/saga"
	"go-link/orchestrator/internal/ports"
)

const (
	PlanIDFree = 1
)

// State holds the state during the Registration Saga execution.
type State struct {
	// Input
	Username  string
	Password  string
	FirstName string
	LastName  string
	Gender    int
	Birthday  string

	// Output
	UserID         int64
	TenantID       int64
	SubscriptionID int64
}

// Saga represents the User Registration Saga.
type Saga struct {
	state          *State
	identityClient ports.IdentityClient
	billingClient  ports.BillingClient
}

// NewSaga creates a new Registration Saga from a RegisterRequest DTO.
func NewSaga(
	req *dto.RegisterRequest,
	identityClient ports.IdentityClient,
	billingClient ports.BillingClient,
) *Saga {
	return &Saga{
		state: &State{
			Username:  req.Username,
			Password:  req.Password,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Gender:    req.Gender,
			Birthday:  req.Birthday,
		},
		identityClient: identityClient,
		billingClient:  billingClient,
	}
}

// BuildSteps returns the ordered steps for this saga.
func (s *Saga) BuildSteps() []saga.Step {
	return []saga.Step{
		&StepCreateUser{State: s.state, IdentityClient: s.identityClient},
		&StepCreateSubscription{State: s.state, BillingClient: s.billingClient, PlanID: PlanIDFree},
	}
}

// Execute runs the Registration Saga and returns the result.
func (s *Saga) Execute(ctx context.Context) (*dto.RegisterResponse, error) {
	coordinator := saga.NewCoordinator(s.BuildSteps()...)

	if err := coordinator.Execute(ctx); err != nil {
		return &dto.RegisterResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &dto.RegisterResponse{
		UserID:   s.state.UserID,
		TenantID: s.state.TenantID,
		Success:  true,
		Message:  "User registered successfully",
	}, nil
}
