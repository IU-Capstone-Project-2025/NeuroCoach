package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if hash == "" {
		t.Error("Expected non-empty hash")
	}

	if hash == password {
		t.Error("Hash should not equal original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword123"
	wrongPassword := "wrongpassword"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Test correct password
	if !CheckPasswordHash(password, hash) {
		t.Error("Expected password to match hash")
	}

	// Test wrong password
	if CheckPasswordHash(wrongPassword, hash) {
		t.Error("Expected wrong password to not match hash")
	}
}

func TestHashPassword_DifferentHashes(t *testing.T) {
	password := "testpassword123"

	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)

	if err1 != nil || err2 != nil {
		t.Errorf("Expected no errors, got %v, %v", err1, err2)
	}

	// Hashes should be different due to salt
	if hash1 == hash2 {
		t.Error("Expected different hashes for same password")
	}

	// But both should validate correctly
	if !CheckPasswordHash(password, hash1) || !CheckPasswordHash(password, hash2) {
		t.Error("Both hashes should validate the password")
	}
}
