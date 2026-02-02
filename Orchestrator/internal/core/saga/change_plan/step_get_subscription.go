package change_plan

import (
	"context"

	"go-link/orchestrator/internal/ports"
)

// StepGetSubscription retrieves current subscription details.
type StepGetSubscription struct {
	State         *State
	BillingClient ports.BillingClient
}

func (s *StepGetSubscription) Name() string { return "GetSubscription" }

func (s *StepGetSubscription) Execute(ctx context.Context) error {
	sub, err := s.BillingClient.GetSubscription(ctx, s.State.SubscriptionID)
	if err != nil {
		return err
	}

	s.State.OldPlanID = sub.PlanID
	// Also ensure TenantID matches
	if s.State.TenantID == 0 {
		s.State.TenantID = sub.TenantID
	}

	return nil
}

func (s *StepGetSubscription) Compensate(ctx context.Context) error {
	return nil
}
