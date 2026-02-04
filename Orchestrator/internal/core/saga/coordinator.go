package saga

import (
	"context"
	"go-link/orchestrator/global"

	"go.uber.org/zap"
)

// Coordinator manages the execution of a Saga
type Coordinator struct {
	steps []Step
}

// NewCoordinator creates a new Saga Coordinator
func NewCoordinator(steps ...Step) *Coordinator {
	return &Coordinator{
		steps: steps,
	}
}

// Execute runs the Saga
func (c *Coordinator) Execute(ctx context.Context) error {
	executedSteps := make([]Step, 0, len(c.steps))

	for _, step := range c.steps {
		global.LoggerZap.Info("Executing Saga Step", zap.String("step", step.Name()))
		if err := step.Execute(ctx); err != nil {
			global.LoggerZap.Warn("Saga Step Failed", zap.String("step", step.Name()), zap.Error(err))
			c.compensate(ctx, executedSteps)
			return err
		}
		executedSteps = append(executedSteps, step)
	}

	global.LoggerZap.Info("Saga Completed Successfully")
	return nil
}

// compensate rolls back executed steps in reverse order
func (c *Coordinator) compensate(ctx context.Context, executedSteps []Step) {
	global.LoggerZap.Info("Starting Compensation")
	for i := len(executedSteps) - 1; i >= 0; i-- {
		step := executedSteps[i]
		global.LoggerZap.Info("Compensating Saga Step", zap.String("step", step.Name()))
		if err := step.Compensate(ctx); err != nil {
			// In production, this should probably retry or alert heavily
			global.LoggerZap.Error("Compensation Failed", zap.String("step", step.Name()), zap.Error(err))
		}
	}
	global.LoggerZap.Info("Compensation Finished")
}
