package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"
	"go-link/identity/internal/ports"
)

// TenantMemberContainer holds tenant member related dependencies.
type TenantMemberContainer struct {
	Repository ports.TenantMemberRepository
}

// InitTenantMemberDependencies initializes tenant member dependencies.
func InitTenantMemberDependencies(client *dbEnt.EntClient) TenantMemberContainer {
	repository := db.NewTenantMemberRepository(client)

	return TenantMemberContainer{
		Repository: repository,
	}
}
