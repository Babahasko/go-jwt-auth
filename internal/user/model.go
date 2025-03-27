package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `gorm:"index" json:"email"`
	Name    string `json:"name"`
	Pasword string `json:"password"`
}
