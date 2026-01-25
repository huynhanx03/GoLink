package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"
	"go-link/identity/internal/ports"
)

// UserAttributeValueContainer holds user attribute value related dependencies.
type UserAttributeValueContainer struct {
	Repository ports.UserAttributeValueRepository
}

// InitUserAttributeValueDependencies initializes user attribute value dependencies.
func InitUserAttributeValueDependencies(client *dbEnt.EntClient) UserAttributeValueContainer {
	repository := db.NewUserAttributeValueRepository(client)

	return UserAttributeValueContainer{
		Repository: repository,
	}
}
