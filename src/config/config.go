package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	IsProduction   bool
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

	isProductionInt, _ := strconv.Atoi(os.Getenv("IS_PRODUCTION"))

	var isProduction bool
	if isProductionInt == 0 {
		isProduction = false
	} else {
		isProduction = true
	}

	return Config{
		Port:         os.Getenv("INTERNAL_PORT"),
		IsProduction: isProduction,
		DatabaseConfig: DatabaseConfig{
			User:      os.Getenv("DB_USER"),
			EntryPort: os.Getenv("DB_ENTRY_PORT"),
			Name:      os.Getenv("DB_NAME"),
			Password:  os.Getenv("DB_PASSWORD"),
		},
	}
}
