package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/ports"
)

// UserAttributeValueContainer holds user attribute value related dependencies.
type UserAttributeValueContainer struct {
	Repository ports.UserAttributeValueRepository
}

// InitUserAttributeValueDependencies initializes user attribute value dependencies.
func InitUserAttributeValueDependencies(client *generate.Client) UserAttributeValueContainer {
	repository := db.NewUserAttributeValueRepository(client)

	return UserAttributeValueContainer{
		Repository: repository,
	}
}
