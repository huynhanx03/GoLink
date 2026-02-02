package change_plan

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

// StepUpdateSubscription updates the subscription to the new plan.
type StepApplyNewPlan struct {
	State         *State
	BillingClient ports.BillingClient
}

func (s *StepApplyNewPlan) Name() string { return "ApplyNewPlan" }

func (s *StepApplyNewPlan) Execute(ctx context.Context) error {
	req := dto.UpdateSubscriptionRequest{
		PlanID: s.State.NewPlanID,
	}
	return s.BillingClient.UpdateSubscription(ctx, s.State.TenantID, req)
}

func (s *StepApplyNewPlan) Compensate(ctx context.Context) error {
	// Revert to old plan
	req := dto.UpdateSubscriptionRequest{
		PlanID: s.State.OldPlanID,
	}
	return s.BillingClient.UpdateSubscription(ctx, s.State.TenantID, req)
}
