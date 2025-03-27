package main

import (
	"log"
	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/internal/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadConfig()
	db, err := gorm.Open(sqlite.Open(conf.DB.DbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	// Migrate the schema
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration success")
}
