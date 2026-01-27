package infrastructure

import (
	"log"

	"go-link/common/pkg/database/ent"

	"go-link/billing/global"
	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
	"go-link/billing/internal/adapters/driven/db/ent/generate"
)

// SetupEnt initializes the Ent client for PostgreSQL.
func SetupEnt() {
	driver, err := ent.NewDriver(global.Config.Database)
	if err != nil {
		log.Fatalf("failed opening connection to ent: %v", err)
	}

	client := generate.NewClient(generate.Driver(driver))

	global.EntClient = dbEnt.WrapClient(client, global.LoggerZap.Logger)
}
