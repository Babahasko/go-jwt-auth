package user

import (
	"github.com/Babahasko/go-jwt-auth/pkg/db"
)

type UserRepository struct{
	Database *db.DB
}

func NewUserRepository(database *db.DB) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

