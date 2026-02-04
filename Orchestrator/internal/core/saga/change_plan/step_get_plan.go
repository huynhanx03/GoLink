package change_plan

import (
	"context"

	"go-link/orchestrator/internal/constant"
	"go-link/orchestrator/internal/ports"
)

type StepGetPlan struct {
	State         *State
	BillingClient ports.BillingClient
}

func (s *StepGetPlan) Name() string { return "GetPlan" }

func (s *StepGetPlan) Execute(ctx context.Context) error {
	plan, err := s.BillingClient.GetPlan(ctx, s.State.NewPlanID)
	if err != nil {
		return err
	}

	s.State.PlanPrice = plan.BasePrice
	s.State.Currency = constant.DefaultCurrency

	return nil
}

func (s *StepGetPlan) Compensate(ctx context.Context) error {
	return nil
}
