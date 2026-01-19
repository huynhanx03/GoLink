package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/ports"
)

// CredentialContainer holds credential-related dependencies.
type CredentialContainer struct {
	Repository ports.CredentialRepository
}

// InitCredentialDependencies initializes credential dependencies.
func InitCredentialDependencies(client *generate.Client) CredentialContainer {
	repository := db.NewCredentialRepository(client)

	return CredentialContainer{
		Repository: repository,
	}
}
