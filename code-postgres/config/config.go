package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Type     string // e.g., "postgres", "inmemory"
	Host     string
	User     string
	Password string
	SSLMode  string
	DBName   string
}

func LoadDatabaseConfig(filename ...string) DatabaseConfig {
	
	// Default to ".env" if no filename is provided
	envFile := ".env"
	if len(filename) > 0 {
		envFile = filename[0]
	}

	// Load the specified .env file
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	return DatabaseConfig{
		Type:     os.Getenv("DBTYPE"),
		Host:     os.Getenv("DBHOST"),
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		SSLMode:  os.Getenv("DBSSLMODE"),
		DBName:   os.Getenv("DBNAME"),
	}
}
