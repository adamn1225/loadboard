package config

import (
	"log"
	"net/url"
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
    serverPortStr := getEnv("SERVER_PORT", "8080")
    serverPort, err := strconv.Atoi(serverPortStr)
    if err != nil {
        log.Printf("Invalid SERVER_PORT (%s), falling back to default: 8080", serverPortStr)
        serverPort = 8080
    }

    databaseURL := getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/loadboard")
    if _, err := url.ParseRequestURI(databaseURL); err != nil {
        log.Fatalf("Invalid DATABASE_URL: %v", err)
    }

    return &Config{
        ServerPort:  serverPort,
        DatabaseURL: databaseURL,
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