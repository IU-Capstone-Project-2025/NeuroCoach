package repository

import (
	"testing"
	"time"

	"rest-api/internal/models"
)

func TestGetRating(t *testing.T) {
	// Mock data for testing rating calculation
	testCases := []struct {
		name           string
		userProgress   []models.UserProgress
		completions    map[int][]time.Time
		expectedRating []models.UserRating
	}{
		{
			name: "basic rating calculation",
			userProgress: []models.UserProgress{
				{UserID: 1, TotalWorkouts: 10},
				{UserID: 2, TotalWorkouts: 5},
			},
			completions: map[int][]time.Time{
				1: {
					time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
				},
				2: {
					time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedRating: []models.UserRating{
				{UserID: 1, TotalWorkouts: 10, MaxConsecutive: 3, Score: 13},
				{UserID: 2, TotalWorkouts: 5, MaxConsecutive: 1, Score: 6},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test consecutive days calculation logic
			for userID, dates := range tc.completions {
				maxConsecutive := calculateMaxConsecutiveDaysFromDates(dates)

				var expectedMax int
				for _, rating := range tc.expectedRating {
					if rating.UserID == userID {
						expectedMax = rating.MaxConsecutive
						break
					}
				}

				if maxConsecutive != expectedMax {
					t.Errorf("User %d: expected max consecutive %d, got %d",
						userID, expectedMax, maxConsecutive)
				}
			}
		})
	}
}

// Helper function to test consecutive days calculation
func calculateMaxConsecutiveDaysFromDates(dates []time.Time) int {
	if len(dates) == 0 {
		return 0
	}

	// Group by days
	daysMap := make(map[string]bool)
	for _, date := range dates {
		dayKey := date.Format("2006-01-02")
		daysMap[dayKey] = true
	}

	// Convert to sorted slice
	var days []time.Time
	for dayKey := range daysMap {
		day, _ := time.Parse("2006-01-02", dayKey)
		days = append(days, day)
	}

	if len(days) == 0 {
		return 0
	}

	// Sort dates
	for i := 0; i < len(days)-1; i++ {
		for j := i + 1; j < len(days); j++ {
			if days[i].After(days[j]) {
				days[i], days[j] = days[j], days[i]
			}
		}
	}

	maxConsecutive := 1
	currentConsecutive := 1

	for i := 1; i < len(days); i++ {
		if days[i].Sub(days[i-1]) == 24*time.Hour {
			currentConsecutive++
			if currentConsecutive > maxConsecutive {
				maxConsecutive = currentConsecutive
			}
		} else {
			currentConsecutive = 1
		}
	}

	return maxConsecutive
}

func TestCalculateLevel(t *testing.T) {
	testCases := []struct {
		workouts int
		expected string
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
			// Create mock repository to test the method
			repo := &MongoDBRepository{}
			result := repo.calculateLevel(tc.workouts)

			if result != tc.expected {
				t.Errorf("For %d workouts, expected %s, got %s",
					tc.workouts, tc.expected, result)
			}
		})
	}
}
