package di

import (
	"github.com/Babahasko/go-jwt-auth/internal/user"
)

type IUserRepository interface{
	Create(user *user.User) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Delete(id uint) error
}