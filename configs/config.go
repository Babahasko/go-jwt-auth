package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	DB DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	DbFile string
}

type AuthConfig struct {
	privateKeyFile string
	publicKeyFile string

}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config {
		DB: DbConfig {
			DbFile: os.Getenv("DB_FILE"),
		},
		Auth: AuthConfig{
			privateKeyFile: os.Getenv("PRIVATE_KEY"),
			publicKeyFile: os.Getenv("PUBLIC_KEY"),
		},
	}
}