package services

import (
	"context"
	"testing"

	"rest-api/internal/models"
)

// Mock repository for testing
type mockMongoDBRepo struct {
	getRatingFunc func(ctx context.Context) ([]models.UserRating, error)
}

func (m *mockMongoDBRepo) GetRating(ctx context.Context) ([]models.UserRating, error) {
	if m.getRatingFunc != nil {
		return m.getRatingFunc(ctx)
	}
	return []models.UserRating{
		{UserID: 1, TotalWorkouts: 15, MaxConsecutive: 7, Score: 22},
		{UserID: 2, TotalWorkouts: 10, MaxConsecutive: 3, Score: 13},
		{UserID: 3, TotalWorkouts: 8, MaxConsecutive: 5, Score: 13},
	}, nil
}

func (m *mockMongoDBRepo) SaveChatMessage(ctx context.Context, message *models.ChatMessage) error {
	return nil
}

func (m *mockMongoDBRepo) GetChatHistory(ctx context.Context, userID int) ([]models.ChatMessage, error) {
	return []models.ChatMessage{}, nil
}

func (m *mockMongoDBRepo) SaveWorkoutPlan(ctx context.Context, plan *models.WorkoutPlan) error {
	return nil
}

func (m *mockMongoDBRepo) GetWorkoutPlan(ctx context.Context, userID int) (*models.WorkoutPlan, error) {
	return nil, nil
}

func (m *mockMongoDBRepo) GetWorkoutByID(ctx context.Context, userID int, workoutID string) (*models.Workout, error) {
	return nil, nil
}

func (m *mockMongoDBRepo) CompleteWorkout(ctx context.Context, userID int, workoutID string) error {
	return nil
}

func (m *mockMongoDBRepo) GetUserProgress(ctx context.Context, userID int) (*models.UserProgress, error) {
	return nil, nil
}

func (m *mockMongoDBRepo) GetShortPlan(ctx context.Context, userID int) (*models.ShortWorkoutPlan, error) {
	return nil, nil
}

func (m *mockMongoDBRepo) SaveShortPlan(ctx context.Context, plan *models.ShortWorkoutPlan) error {
	return nil
}

func TestAIService_GetRating(t *testing.T) {
	mockRepo := &mockMongoDBRepo{}
	service := &AIService{
		BaseService: BaseService{MongoDBRepo: mockRepo},
	}

	ctx := context.Background()
	ratings, err := service.GetRating(ctx)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(ratings) != 3 {
		t.Errorf("Expected 3 ratings, got %d", len(ratings))
	}

	// Check if ratings are sorted by score (descending)
	if ratings[0].Score < ratings[1].Score {
		t.Errorf("Ratings should be sorted by score descending")
	}

	// Check score calculation
	expectedScore := ratings[0].TotalWorkouts + ratings[0].MaxConsecutive
	if ratings[0].Score != expectedScore {
		t.Errorf("Score calculation incorrect: expected %d, got %d",
			expectedScore, ratings[0].Score)
	}
}

func TestAIService_GetRating_Error(t *testing.T) {
	mockRepo := &mockMongoDBRepo{
		getRatingFunc: func(ctx context.Context) ([]models.UserRating, error) {
			return nil, NewServiceError(500, "Database error", nil)
		},
	}

	service := &AIService{
		BaseService: BaseService{MongoDBRepo: mockRepo},
	}

	ctx := context.Background()
	_, err := service.GetRating(ctx)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
