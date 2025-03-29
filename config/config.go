package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds the configuration values for the Loadboard project.
type Config struct {
	ServerPort  int
	DatabaseURL string
	Environment string
}

// LoadConfig initializes and returns a Config struct populated with environment variables.
func LoadConfig() *Config {
	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		log.Fatalf("Invalid SERVER_PORT: %v", err)
	}

	return &Config{
		ServerPort:  serverPort,
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/loadboard"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv retrieves the value of the environment variable named by the key or returns the fallback value if not set.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
