package utils

import (
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	userID := 123
	email := "test@example.com"
	secret := "test-secret"
	expiration := time.Hour

	token, err := GenerateJWT(userID, email, secret, expiration)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Error("Expected non-empty token")
	}
}

func TestValidateJWT(t *testing.T) {
	userID := 123
	email := "test@example.com"
	secret := "test-secret"
	expiration := time.Hour

	// Generate token
	token, err := GenerateJWT(userID, email, secret, expiration)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate token
	claims, err := ValidateJWT(token, secret)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if claims == nil {
		t.Error("Expected non-nil claims")
	}

	// Check claims
	if claims["email"] != email {
		t.Errorf("Expected email %s, got %v", email, claims["email"])
	}
}

func TestValidateJWT_InvalidSecret(t *testing.T) {
	userID := 123
	email := "test@example.com"
	secret := "test-secret"
	wrongSecret := "wrong-secret"
	expiration := time.Hour

	// Generate token with correct secret
	token, err := GenerateJWT(userID, email, secret, expiration)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Try to validate with wrong secret
	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Error("Expected error for wrong secret")
	}
}

func TestGetUserIDFromClaims(t *testing.T) {
	userID := 123
	email := "test@example.com"
	secret := "test-secret"
	expiration := time.Hour

	// Generate and validate token
	token, _ := GenerateJWT(userID, email, secret, expiration)
	claims, _ := ValidateJWT(token, secret)

	// Extract user ID
	extractedUserID, err := GetUserIDFromClaims(claims)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if extractedUserID != userID {
		t.Errorf("Expected user ID %d, got %d", userID, extractedUserID)
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	userID := 123
	email := "test@example.com"
	secret := "test-secret"
	expiration := -time.Hour // Expired token

	// Generate expired token
	token, err := GenerateJWT(userID, email, secret, expiration)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Try to validate expired token
	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Error("Expected error for expired token")
	}
}
