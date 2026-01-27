package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"
	"go-link/identity/internal/ports"
)

// CredentialContainer holds credential-related dependencies.
type CredentialContainer struct {
	Repository ports.CredentialRepository
}

// InitCredentialDependencies initializes credential dependencies.
func InitCredentialDependencies(client *dbEnt.EntClient) CredentialContainer {
	repository := db.NewCredentialRepository(client)

	return CredentialContainer{
		Repository: repository,
	}
}
