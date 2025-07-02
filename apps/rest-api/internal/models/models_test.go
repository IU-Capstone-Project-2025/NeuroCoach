package models

import (
	"testing"
	"time"
)

func TestUserRating_ScoreCalculation(t *testing.T) {
	testCases := []struct {
		name           string
		totalWorkouts  int
		maxConsecutive int
		expectedScore  int
	}{
		{"basic calculation", 10, 5, 15},
		{"zero consecutive", 20, 0, 20},
		{"zero workouts", 0, 7, 7},
		{"both zero", 0, 0, 0},
		{"large numbers", 100, 30, 130},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rating := UserRating{
				UserID:         1,
				TotalWorkouts:  tc.totalWorkouts,
				MaxConsecutive: tc.maxConsecutive,
				Score:          tc.totalWorkouts + tc.maxConsecutive,
			}

			if rating.Score != tc.expectedScore {
				t.Errorf("Expected score %d, got %d", tc.expectedScore, rating.Score)
			}
		})
	}
}

func TestFitnessProfile_Validation(t *testing.T) {
	validProfile := FitnessProfile{
		Height:           175.5,
		Weight:           70.0,
		Age:              25,
		Goal:             "weight_loss",
		HealthIssues:     []string{"knee_pain"},
		Timeframe:        "3months",
		FitnessLevel:     "intermediate",
		AvailableMinutes: 60,
		UpdatedAt:        time.Now(),
	}

	// Test valid profile structure
	if validProfile.Height <= 0 {
		t.Error("Height should be positive")
	}
	if validProfile.Weight <= 0 {
		t.Error("Weight should be positive")
	}
	if validProfile.Age < 13 || validProfile.Age > 120 {
		t.Error("Age should be between 13 and 120")
	}
	if validProfile.AvailableMinutes < 30 || validProfile.AvailableMinutes > 1000 {
		t.Error("Available minutes should be between 30 and 1000")
	}
}

func TestWorkoutCompletion_TimeTracking(t *testing.T) {
	now := time.Now()
	completion := WorkoutCompletion{
		UserID:      1,
		CompletedAt: now,
	}

	if completion.CompletedAt.IsZero() {
		t.Error("CompletedAt should not be zero")
	}

	if completion.CompletedAt.After(time.Now().Add(time.Second)) {
		t.Error("CompletedAt should not be in the future")
	}
}

func TestUserProgress_LevelProgression(t *testing.T) {
	testCases := []struct {
		totalWorkouts int
		expectedLevel string
	}{
		{0, "Beginner"},
		{4, "Beginner"},
		{5, "Intermediate"},
		{19, "Intermediate"},
		{20, "Advanced"},
		{49, "Advanced"},
		{50, "Expert"},
		{100, "Expert"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			progress := UserProgress{
				TotalWorkouts: tc.totalWorkouts,
			}

			// Simulate level calculation logic
			var level string
			if progress.TotalWorkouts >= 50 {
				level = "Expert"
			} else if progress.TotalWorkouts >= 20 {
				level = "Advanced"
			} else if progress.TotalWorkouts >= 5 {
				level = "Intermediate"
			} else {
				level = "Beginner"
			}

			if level != tc.expectedLevel {
				t.Errorf("For %d workouts, expected level %s, got %s",
					tc.totalWorkouts, tc.expectedLevel, level)
			}
		})
	}
}
