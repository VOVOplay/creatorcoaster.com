package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	User      string
	EntryPort string
	Name      string
	Password  string
}

func GetConfig() Config {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Could not load .env file. Crashing.")
	}

	return Config{
		Port: os.Getenv("INTERNAL_PORT"),
		DatabaseConfig: DatabaseConfig{
			User:      os.Getenv("DB_USER"),
			EntryPort: os.Getenv("DB_ENTRY_PORT"),
			Name:      os.Getenv("DB_NAME"),
			Password:  os.Getenv("DB_PASSWORD"),
		},
	}
}
