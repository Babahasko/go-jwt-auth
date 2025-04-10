package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `gorm:"index" json:"email"`
	Name    string `json:"name"`
	Password string `json:"password"`
}
