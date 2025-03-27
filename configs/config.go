package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	DB DbConfig
}

type DbConfig struct {
	DbFile string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config {
		DB: DbConfig {
			DbFile: os.Getenv("DbFile"),
		},
	}
}