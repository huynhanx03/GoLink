package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/ports"
)

// TenantMemberContainer holds tenant member related dependencies.
type TenantMemberContainer struct {
	Repository ports.TenantMemberRepository
}

// InitTenantMemberDependencies initializes tenant member dependencies.
func InitTenantMemberDependencies(client *generate.Client) TenantMemberContainer {
	repository := db.NewTenantMemberRepository(client)

	return TenantMemberContainer{
		Repository: repository,
	}
}
