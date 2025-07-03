package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"rest-api/internal/config"
	"rest-api/internal/handlers"
	"rest-api/internal/models"
	"rest-api/internal/services"

	"github.com/gorilla/mux"
)

// Integration test for the rating endpoint
func TestRatingEndpointIntegration(t *testing.T) {
	// Skip if no test database available
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Setup test configuration
	cfg := &config.Config{
		DatabaseURL:   "postgres://postgres:postgres@localhost:5432/fitness_ai_test?sslmode=disable",
		MongoURI:      "mongodb://localhost:27017/fitness_ai_test",
		MongoDBName:   "fitness_ai_test",
		JWTSecret:     "test-secret",
		JWTExpiration: time.Hour,
	}

	// Initialize repositories (would need actual test DB)
	// This is a mock setup for demonstration
	mockPostgresRepo := &mockPostgresRepo{}
	mockMongoRepo := &mockMongoRepo{}

	// Initialize services
	authService := services.NewAuthService(mockPostgresRepo, cfg.JWTSecret, cfg.JWTExpiration)
	profileService := services.NewProfileService(mockPostgresRepo)
	aiService := services.NewAIService(mockPostgresRepo, mockMongoRepo, "")
	healthService := services.NewHealthService(mockPostgresRepo)

	// Initialize handlers
	h := handlers.NewHandlers(authService, profileService, aiService, healthService)

	// Setup router
	r := mux.NewRouter()
	authRouter := r.PathPrefix("/api").Subrouter()
	authRouter.HandleFunc("/rating", h.GetRating).Methods("GET")

	// Create test request
	req := httptest.NewRequest("GET", "/api/rating", nil)
	w := httptest.NewRecorder()

	// Execute request
	r.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var ratings []models.UserRating
	if err := json.Unmarshal(w.Body.Bytes(), &ratings); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(ratings) == 0 {
		t.Error("Expected at least one rating")
	}
}

// Mock repositories for integration testing
type mockPostgresRepo struct{}

func (m *mockPostgresRepo) CreateUser(ctx context.Context, email, passwordHash string) (int, error) {
	return 1, nil
}

func (m *mockPostgresRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{ID: 1, Email: email}, nil
}

func (m *mockPostgresRepo) SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error {
	return nil
}

func (m *mockPostgresRepo) GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error) {
	return &models.FitnessProfile{}, nil
}

func (m *mockPostgresRepo) Ping(ctx context.Context) error {
	return nil
}

func (m *mockPostgresRepo) Close() error {
	return nil
}

type mockMongoRepo struct{}

func (m *mockMongoRepo) SaveChatMessage(ctx context.Context, message *models.ChatMessage) error {
	return nil
}

func (m *mockMongoRepo) GetChatHistory(ctx context.Context, userID int) ([]models.ChatMessage, error) {
	return []models.ChatMessage{}, nil
}

func (m *mockMongoRepo) SaveWorkoutPlan(ctx context.Context, plan *models.WorkoutPlan) error {
	return nil
}

func (m *mockMongoRepo) GetWorkoutPlan(ctx context.Context, userID int) (*models.WorkoutPlan, error) {
	return nil, nil
}

func (m *mockMongoRepo) GetWorkoutByID(ctx context.Context, userID int, workoutID string) (*models.Workout, error) {
	return nil, nil
}

func (m *mockMongoRepo) CompleteWorkout(ctx context.Context, userID int, workoutID string) error {
	return nil
}

func (m *mockMongoRepo) GetUserProgress(ctx context.Context, userID int) (*models.UserProgress, error) {
	return &models.UserProgress{}, nil
}

func (m *mockMongoRepo) GetShortPlan(ctx context.Context, userID int) (*models.ShortWorkoutPlan, error) {
	return nil, nil
}

func (m *mockMongoRepo) SaveShortPlan(ctx context.Context, plan *models.ShortWorkoutPlan) error {
	return nil
}

func (m *mockMongoRepo) GetRating(ctx context.Context) ([]models.UserRating, error) {
	return []models.UserRating{
		{UserID: 1, TotalWorkouts: 25, MaxConsecutive: 7, Score: 32},
		{UserID: 2, TotalWorkouts: 15, MaxConsecutive: 5, Score: 20},
	}, nil
}

// Benchmark test for rating calculation
func BenchmarkRatingCalculation(b *testing.B) {
	// Sample data for benchmarking
	dates := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateMaxConsecutiveDaysFromDates(dates)
	}
}

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

// Example of how to run the integration test
func ExampleRatingEndpoint() {
	// This would be used in actual integration testing
	// with real database connections
	fmt.Println("Rating endpoint integration test example")
	// Output: Rating endpoint integration test example
}
