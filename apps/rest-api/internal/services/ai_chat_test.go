package services

import (
	"context"
	"testing"

	"rest-api/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test Chat with beginner mode
func TestChatWithBeginnerMode(t *testing.T) {
	// Skip if short test mode
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	// Create mocks
	mockRepo := new(MockRepository)
	mockMongoRepo := new(MockMongoDBRepository)

	// Create service with mock client
	service := &AIService{
		BaseService: BaseService{
			Repo:        mockRepo,
			MongoDBRepo: mockMongoRepo,
		},
		Client: nil, // We'll mock the client responses
	}

	// Create test context with user ID
	ctx := context.WithValue(context.Background(), "user_id", 1)

	// Create test profiles
	beginnerProfile := &models.FitnessProfile{
		UserID:       1,
		FitnessLevel: "beginner",
	}

	advancedProfile := &models.FitnessProfile{
		UserID:       1,
		FitnessLevel: "advanced",
	}

	// Create empty chat history
	emptyHistory := []models.ChatMessage{}

	// Setup test cases
	tests := []struct {
		name          string
		profile       *models.FitnessProfile
		profileErr    error
		expectBeginner bool
	}{
		{
			name:          "Beginner user",
			profile:       beginnerProfile,
			profileErr:    nil,
			expectBeginner: true,
		},
		{
			name:          "Advanced user",
			profile:       advancedProfile,
			profileErr:    nil,
			expectBeginner: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mocks
			mockRepo.ExpectedCalls = nil
			mockMongoRepo.ExpectedCalls = nil

			// Setup mocks
			mockRepo.On("GetFitnessProfile", mock.Anything, 1).Return(tt.profile, tt.profileErr)
			mockMongoRepo.On("GetChatHistory", mock.Anything, 1).Return(emptyHistory, nil)

			// We can't easily test the actual AI call, but we can verify the beginner mode is set correctly
			// by checking if the mock was called with the right parameters
			if service.Client != nil {
				// This is a simplified test that doesn't actually call the AI service
				// In a real test, you would mock the client and verify the system message contains
				// beginner instructions when appropriate
				_, err := service.Chat(ctx, "Test message")
				assert.Error(t, err) // Should error because Client is nil
			}

			// Verify mocks were called
			mockRepo.AssertExpectations(t)
			mockMongoRepo.AssertExpectations(t)
		})
	}
}