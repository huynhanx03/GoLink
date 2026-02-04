package change_plan

import (
	"context"

	"go-link/orchestrator/internal/ports"
)

type StepUpdateTenantPlan struct {
	State          *State
	identityClient ports.IdentityClient
}

func NewStepUpdateTenantPlan(state *State, identityClient ports.IdentityClient) *StepUpdateTenantPlan {
	return &StepUpdateTenantPlan{
		State:          state,
		identityClient: identityClient,
	}
}

func (s *StepUpdateTenantPlan) Name() string {
	return "UpdateTenantPlan"
}

func (s *StepUpdateTenantPlan) Execute(ctx context.Context) error {
	return s.identityClient.UpdateTenantPlan(ctx, s.State.TenantID, int64(s.State.NewPlanID))
}

func (s *StepUpdateTenantPlan) Compensate(ctx context.Context) error {
	if s.State.OldPlanID > 0 {
		return s.identityClient.UpdateTenantPlan(ctx, s.State.TenantID, int64(s.State.OldPlanID))
	}

	return nil
}
