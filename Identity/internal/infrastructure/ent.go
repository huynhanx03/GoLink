package infrastructure

import (
	"log"

	"go-link/common/pkg/database/ent"

	"go-link/identity/global"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
)

// SetupEnt initializes the Ent client for PostgreSQL.
func SetupEnt() {
	driver, err := ent.NewDriver(global.Config.Database)
	if err != nil {
		log.Fatalf("failed opening connection to ent: %v", err)
	}

	client := generate.NewClient(generate.Driver(driver))
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	global.EntClient = client
}
