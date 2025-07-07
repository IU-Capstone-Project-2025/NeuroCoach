package config

import (
	"os"
	"testing"
	"time"
)

func TestLoad_DefaultValues(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("PORT")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("JWT_EXPIRATION")
	os.Unsetenv("OPENROUTER_KEY")
	os.Unsetenv("MONGOURI")
	os.Unsetenv("MONGODBNAME")
	os.Unsetenv("SKIP_DATABASE")

	cfg, err := Load()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if cfg.Environment != "development" {
		t.Errorf("Expected environment 'development', got %s", cfg.Environment)
	}

	if cfg.Port != "8080" {
		t.Errorf("Expected port '8080', got %s", cfg.Port)
	}

	if cfg.JWTExpiration != 24*time.Hour {
		t.Errorf("Expected JWT expiration 24h, got %v", cfg.JWTExpiration)
	}

	if cfg.SkipDatabase {
		t.Error("Expected SkipDatabase to be false by default")
	}
}

func TestLoad_EnvironmentVariables(t *testing.T) {
	// Set environment variables
	os.Setenv("ENVIRONMENT", "production")
	os.Setenv("PORT", "3000")
	os.Setenv("JWT_SECRET", "custom-secret")
	os.Setenv("JWT_EXPIRATION", "12h")
	os.Setenv("MONGOURI", "mongodb://custom:27017")
	os.Setenv("MONGODBNAME", "custom_db")
	os.Setenv("SKIP_DATABASE", "true")

	defer func() {
		// Clean up
		os.Unsetenv("ENVIRONMENT")
		os.Unsetenv("PORT")
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("JWT_EXPIRATION")
		os.Unsetenv("MONGOURI")
		os.Unsetenv("MONGODBNAME")
		os.Unsetenv("SKIP_DATABASE")
	}()

	cfg, err := Load()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if cfg.Environment != "production" {
		t.Errorf("Expected environment 'production', got %s", cfg.Environment)
	}

	if cfg.Port != "3000" {
		t.Errorf("Expected port '3000', got %s", cfg.Port)
	}

	if cfg.JWTSecret != "custom-secret" {
		t.Errorf("Expected JWT secret 'custom-secret', got %s", cfg.JWTSecret)
	}

	if cfg.JWTExpiration != 12*time.Hour {
		t.Errorf("Expected JWT expiration 12h, got %v", cfg.JWTExpiration)
	}

	if cfg.MongoURI != "mongodb://custom:27017" {
		t.Errorf("Expected MongoURI 'mongodb://custom:27017', got %s", cfg.MongoURI)
	}

	if cfg.MongoDBName != "custom_db" {
		t.Errorf("Expected MongoDBName 'custom_db', got %s", cfg.MongoDBName)
	}

	if !cfg.SkipDatabase {
		t.Error("Expected SkipDatabase to be true")
	}
}

func TestGetEnv(t *testing.T) {
	// Test with existing environment variable
	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")

	result := getEnv("TEST_VAR", "fallback")
	if result != "test_value" {
		t.Errorf("Expected 'test_value', got %s", result)
	}

	// Test with non-existing environment variable
	result = getEnv("NON_EXISTING_VAR", "fallback")
	if result != "fallback" {
		t.Errorf("Expected 'fallback', got %s", result)
	}
}

func TestParseDuration_Valid(t *testing.T) {
	testCases := []struct {
		input    string
		expected time.Duration
	}{
		{"1h", time.Hour},
		{"30m", 30 * time.Minute},
		{"24h", 24 * time.Hour},
		{"1h30m", time.Hour + 30*time.Minute},
	}

	for _, tc := range testCases {
		result := parseDuration(tc.input)
		if result != tc.expected {
			t.Errorf("For input %s, expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestParseDuration_Invalid(t *testing.T) {
	// Test with invalid duration format
	result := parseDuration("invalid")
	expected := 24 * time.Hour // Default fallback

	if result != expected {
		t.Errorf("Expected fallback duration %v, got %v", expected, result)
	}
}

func TestLoad_EmptyOpenRouterKey(t *testing.T) {
	// Set empty OpenRouter key
	os.Setenv("OPENROUTER_KEY", "")
	defer os.Unsetenv("OPENROUTER_KEY")

	cfg, err := Load()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if cfg.OpenRouterKey != "" {
		t.Errorf("Expected empty OpenRouter key, got %s", cfg.OpenRouterKey)
	}
}
