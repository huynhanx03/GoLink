package change_plan

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/core/saga"
	"go-link/orchestrator/internal/ports"
)

// Saga represents the Upgrade Subscription Saga.
type Saga struct {
	state          *State
	billingClient  ports.BillingClient
	paymentClient  ports.PaymentClient
	identityClient ports.IdentityClient
}

// NewSaga creates a new Upgrade Subscription Saga.
func NewSaga(
	req *dto.UpgradeSubscriptionRequest,
	billingClient ports.BillingClient,
	paymentClient ports.PaymentClient,
	identityClient ports.IdentityClient,
) *Saga {
	return &Saga{
		state: &State{
			TenantID:       req.TenantID,
			SubscriptionID: req.SubscriptionID,
			NewPlanID:      req.NewPlanID,
		},
		billingClient:  billingClient,
		paymentClient:  paymentClient,
		identityClient: identityClient,
	}
}

// BuildSteps returns the ordered steps for this saga.
func (s *Saga) BuildSteps() []saga.Step {
	return []saga.Step{
		&StepGetSubscription{State: s.state, BillingClient: s.billingClient},
		&StepGetPlan{State: s.state, BillingClient: s.billingClient},
		&StepCreateInvoice{State: s.state, BillingClient: s.billingClient},
		&StepProcessPayment{State: s.state, PaymentClient: s.paymentClient},
		&StepMarkInvoicePaid{State: s.state, BillingClient: s.billingClient},
		&StepApplyNewPlan{State: s.state, BillingClient: s.billingClient},
		&StepUpdateTenantPlan{State: s.state, identityClient: s.identityClient},
	}
}

// Execute runs the Upgrade Subscription Saga.
func (s *Saga) Execute(ctx context.Context) (*dto.UpgradeSubscriptionResponse, error) {
	coordinator := saga.NewCoordinator(s.BuildSteps()...)

	if err := coordinator.Execute(ctx); err != nil {
		return nil, err
	}

	return &dto.UpgradeSubscriptionResponse{
		InvoiceID: s.state.InvoiceID,
	}, nil
}
