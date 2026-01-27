package infrastructure

import (
	"go-link/payment/internal/di"
)

func Run() error {
	LoadConfig()
	SetupLogger()
	di.SetupDependencies()

	return nil
}
