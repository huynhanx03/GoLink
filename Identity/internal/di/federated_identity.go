package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/ports"
)

// FederatedIdentityContainer holds federated identity related dependencies.
type FederatedIdentityContainer struct {
	Repository ports.FederatedIdentityRepository
}

// InitFederatedIdentityDependencies initializes federated identity dependencies.
func InitFederatedIdentityDependencies(client *generate.Client) FederatedIdentityContainer {
	repository := db.NewFederatedIdentityRepository(client)

	return FederatedIdentityContainer{
		Repository: repository,
	}
}
