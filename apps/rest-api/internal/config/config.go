package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment   string
	Port          string
	DatabaseURL   string
	JWTSecret     string
	JWTExpiration time.Duration
	OpenAIKey     string
	MongoURI      string
	MongoDBName   string
}

func Load() (*Config, error) {
	// Load .env file if it exists (for local development)
	_ = godotenv.Load()

	// Set default values
	cfg := &Config{
		Environment:   getEnv("ENVIRONMENT", "development"),
		Port:          getEnv("PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/fitness_ai?sslmode=disable"),
		JWTSecret:     getEnv("JWT_SECRET", "default-secret-change-me"),
		JWTExpiration: parseDuration(getEnv("JWT_EXPIRATION", "24h")),
		OpenAIKey:     getEnv("OPENAI_KEY", ""),
		MongoURI:      getEnv("MONGOURI", "mongodb://user:password@localhost:27017"),
		MongoDBName:   getEnv("MONGODBNAME", "fitness_ai"),
	}

	// Validate required fields
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	if cfg.OpenAIKey == "" {
		log.Println("WARNING: OPENAI_KEY is not set - AI features will be disabled")
	}

	return cfg, nil
}

// Helper function to read environment variables with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper function to parse duration from env
func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Printf("Invalid duration format '%s', defaulting to 24h", durationStr)
		return 24 * time.Hour
	}
	return duration
}
