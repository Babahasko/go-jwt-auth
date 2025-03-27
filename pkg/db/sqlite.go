package db

import (
	"github.com/Babahasko/go-jwt-auth/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *DB {
	db, err := gorm.Open(sqlite.Open(conf.DB.DbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DB{db}
}
