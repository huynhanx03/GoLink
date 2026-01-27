package registration

import (
	"context"

	"go-link/orchestrator/internal/ports"
)

// StepCreateSubscription implements the Billing subscription creation step.
type StepCreateSubscription struct {
	State         *State
	BillingClient ports.BillingClient
	PlanID        int
}

func (s *StepCreateSubscription) Name() string { return "CreateSubscription" }

func (s *StepCreateSubscription) Execute(ctx context.Context) error {
	subID, err := s.BillingClient.CreateSubscription(ctx, s.State.UserID, s.PlanID)
	if err != nil {
		return err
	}
	s.State.SubscriptionID = subID
	return nil
}

func (s *StepCreateSubscription) Compensate(ctx context.Context) error {
	if s.State.SubscriptionID == 0 {
		return nil
	}
	return s.BillingClient.CancelSubscription(ctx, s.State.SubscriptionID)
}
