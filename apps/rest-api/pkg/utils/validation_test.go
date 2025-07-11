package utils

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name@domain.co.uk", true},
		{"invalid-email", false},
		{"@domain.com", false},
		{"user@", false},
		{"", false},
	}

	for _, tc := range testCases {
		result := isValidEmail(tc.email)
		if result != tc.expected {
			t.Errorf("For email '%s', expected %v, got %v", tc.email, tc.expected, result)
		}
	}
}

func TestValidateAge(t *testing.T) {
	testCases := []struct {
		age      int
		expected bool
	}{
		{25, true},
		{18, true},
		{65, true},
		{17, false},
		{0, false},
		{-5, false},
		{150, false},
	}

	for _, tc := range testCases {
		result := isValidAge(tc.age)
		if result != tc.expected {
			t.Errorf("For age %d, expected %v, got %v", tc.age, tc.expected, result)
		}
	}
}

func TestValidateTimeframe(t *testing.T) {
	validTimeframes := []string{"1month", "3months", "6months", "1year"}
	invalidTimeframes := []string{"2months", "5months", "2years", ""}

	for _, timeframe := range validTimeframes {
		if !isValidTimeframe(timeframe) {
			t.Errorf("Expected timeframe '%s' to be valid", timeframe)
		}
	}

	for _, timeframe := range invalidTimeframes {
		if isValidTimeframe(timeframe) {
			t.Errorf("Expected timeframe '%s' to be invalid", timeframe)
		}
	}
}

func TestValidateFitnessLevel(t *testing.T) {
	validLevels := []string{"beginner", "intermediate", "advanced"}
	invalidLevels := []string{"expert", "novice", ""}

	for _, level := range validLevels {
		if !isValidFitnessLevel(level) {
			t.Errorf("Expected fitness level '%s' to be valid", level)
		}
	}

	for _, level := range invalidLevels {
		if isValidFitnessLevel(level) {
			t.Errorf("Expected fitness level '%s' to be invalid", level)
		}
	}
}

// Helper functions for validation
func isValidEmail(email string) bool {
	if len(email) == 0 || len(email) > 254 {
		return false
	}
	// Simple check: must contain @ and have text before and after it
	atIndex := -1
	for i, c := range email {
		if c == '@' {
			if atIndex != -1 { // Multiple @ symbols
				return false
			}
			atIndex = i
		}
	}
	return atIndex > 0 && atIndex < len(email)-1
}

func isValidAge(age int) bool {
	return age >= 18 && age <= 120
}

func isValidTimeframe(timeframe string) bool {
	validTimeframes := map[string]bool{
		"1month":  true,
		"3months": true,
		"6months": true,
		"1year":   true,
	}
	return validTimeframes[timeframe]
}

func isValidFitnessLevel(level string) bool {
	validLevels := map[string]bool{
		"beginner":     true,
		"intermediate": true,
		"advanced":     true,
	}
	return validLevels[level]
}
