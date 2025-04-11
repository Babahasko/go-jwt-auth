package auth

import (
	"errors"

	"github.com/Babahasko/go-jwt-auth/internal/user"
	"github.com/Babahasko/go-jwt-auth/pkg/di"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, err := service.UserRepository.GetByEmail(email)
	
	if err != nil {
		return "", err
	}

	if existedUser != nil {
		return "", errors.New(ErrorUserExist)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	if err != nil {
		return "", err
	}
	
	user := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	result, err := service.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	return result.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, err := service.UserRepository.GetByEmail(email)
	
	if err != nil {
		return "", err
	}

	if existedUser == nil {
		return "", errors.New(ErrorUserNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrorWrongPassword)
	}

	return existedUser.Email, nil
}
