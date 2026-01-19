package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/ports"
)

// UserContainer holds user-related dependencies.
type UserContainer struct {
	Repository ports.UserRepository
}

// InitUserDependencies initializes user dependencies.
func InitUserDependencies(client *generate.Client) UserContainer {
	repository := db.NewUserRepository(client)

	return UserContainer{
		Repository: repository,
	}
}
