package auth

import (
	"github.com/Babahasko/go-jwt-auth/pkg/di"
)

type AuthService struct {
	UserRepository *di.IUserRepository
}

func NewUserService(userRepository *di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}