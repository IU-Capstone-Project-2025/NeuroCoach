package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment       string
	Port              string
	DatabaseURL       string
	JWTSecret         string
	JWTExpiration     time.Duration
	RefreshExpiration time.Duration
	OpenRouterKey     string
	MongoURI          string
	MongoDBName       string
	SkipDatabase      bool
}

func Load() (*Config, error) {
	// Load .env file if it exists (for local development)
	_ = godotenv.Load()

	// Set default values
	cfg := &Config{
		Environment:       getEnv("ENVIRONMENT", "development"),
		Port:              getEnv("PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/fitness_ai?sslmode=disable"),
		JWTSecret:         getEnv("JWT_SECRET", "default-secret-change-me"),
		JWTExpiration:     parseDuration(getEnv("JWT_EXPIRATION", "15m")),
		RefreshExpiration: parseDuration(getEnv("REFRESH_EXPIRATION", "7d")),
		OpenRouterKey:     getEnv("OPENROUTER_KEY", "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"),
		MongoURI:          getEnv("MONGOURI", "mongodb://localhost:27017/fitness_ai"),
		MongoDBName:       getEnv("MONGODBNAME", "fitness_ai"),
		SkipDatabase:      getEnv("SKIP_DATABASE", "") != "",
	}

	// Validate required fields
	if !cfg.SkipDatabase && cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	if cfg.OpenRouterKey == "" {
		log.Println("WARNING: OPENROUTER_KEY is not set - AI features will be disabled")
	}

	if cfg.SkipDatabase {
		log.Println("INFO: Running in database-free mode for AI testing")
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
