package user

import (
	"errors"

	"github.com/Babahasko/go-jwt-auth/pkg/db"
	"gorm.io/gorm"
)

type UserRepository struct{
	Database *db.DB
}

func NewUserRepository(database *db.DB) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //handle not found error
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //handle not found error
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&User{},id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}