package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// The main configuration from enviromental variables.
// Also includes the database configuration.
type Config struct {
	Port           string // The port which the program should be served on
	IsProduction   bool   // Is this prod?
	DatabaseConfig DatabaseConfig
}

// The configuration used to connect to the database
type DatabaseConfig struct {
	User     string
	Address  string
	Name     string
	Password string
}

// Gets all of the config from enviromental variables.
// This includes the database configuration as well.
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
			User:     os.Getenv("DB_USER"),
			Address:  os.Getenv("DB_ADDRESS"),
			Name:     os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}
