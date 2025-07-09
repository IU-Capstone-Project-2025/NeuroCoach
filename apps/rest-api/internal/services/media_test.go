package services

import (
	"context"
	"errors"
	"testing"

	"rest-api/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSaveExerciseMedia(t *testing.T) {
	// Create mocks
	mockRepo := new(MockRepository)
	mockMongoRepo := new(MockMongoDBRepository)
	
	// Create service
	service := NewMediaService(mockRepo, mockMongoRepo)
	
	// Create test context with user ID
	ctx := context.WithValue(context.Background(), "user_id", 1)
	
	// Create valid exercise ID
	validID := primitive.NewObjectID().Hex()
	
	// Setup test cases
	tests := []struct {
		name        string
		request     *models.ExerciseMediaRequest
		setupMocks  func()
		expectError bool
	}{
		{
			name: "Valid request",
			request: &models.ExerciseMediaRequest{
				ExerciseID:  validID,
				ImageURL:    "https://example.com/image.jpg",
				Description: "Test description",
				Order:       1,
			},
			setupMocks: func() {
				mockMongoRepo.On("SaveExerciseMedia", mock.Anything, mock.AnythingOfType("*models.ExerciseMedia")).Return(nil)
			},
			expectError: false,
		},
		{
			name: "Invalid exercise ID",
			request: &models.ExerciseMediaRequest{
				ExerciseID:  "invalid-id",
				ImageURL:    "https://example.com/image.jpg",
				Description: "Test description",
				Order:       1,
			},
			setupMocks:  func() {},
			expectError: true,
		},
		{
			name: "Database error",
			request: &models.ExerciseMediaRequest{
				ExerciseID:  validID,
				ImageURL:    "https://example.com/image.jpg",
				Description: "Test description",
				Order:       1,
			},
			setupMocks: func() {
				mockMongoRepo.On("SaveExerciseMedia", mock.Anything, mock.AnythingOfType("*models.ExerciseMedia")).Return(errors.New("database error"))
			},
			expectError: true,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockMongoRepo.ExpectedCalls = nil
			tt.setupMocks()
			
			// Call the method
			err := service.SaveExerciseMedia(ctx, tt.request)
			
			// Check result
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			
			// Verify mocks
			mockMongoRepo.AssertExpectations(t)
		})
	}
}

func TestGetExerciseMedia(t *testing.T) {
	// Create mocks
	mockRepo := new(MockRepository)
	mockMongoRepo := new(MockMongoDBRepository)
	
	// Create service
	service := NewMediaService(mockRepo, mockMongoRepo)
	
	// Create test context with user ID
	ctx := context.WithValue(context.Background(), "user_id", 1)
	
	// Create valid exercise ID
	validID := primitive.NewObjectID().Hex()
	
	// Create sample media
	sampleMedia := []models.ExerciseMedia{
		{
			ID:          primitive.NewObjectID(),
			ExerciseID:  primitive.ObjectID{},
			ImageURL:    "https://example.com/image1.jpg",
			Description: "Test description 1",
			Order:       1,
		},
		{
			ID:          primitive.NewObjectID(),
			ExerciseID:  primitive.ObjectID{},
			ImageURL:    "https://example.com/image2.jpg",
			Description: "Test description 2",
			Order:       2,
		},
	}
	
	// Setup test cases
	tests := []struct {
		name        string
		exerciseID  string
		setupMocks  func()
		expectError bool
		expected    []models.ExerciseMedia
	}{
		{
			name:       "Valid request",
			exerciseID: validID,
			setupMocks: func() {
				mockMongoRepo.On("GetExerciseMedia", mock.Anything, validID).Return(sampleMedia, nil)
			},
			expectError: false,
			expected:    sampleMedia,
		},
		{
			name:       "Invalid exercise ID",
			exerciseID: "invalid-id",
			setupMocks: func() {},
			expectError: true,
			expected:    nil,
		},
		{
			name:       "Database error",
			exerciseID: validID,
			setupMocks: func() {
				mockMongoRepo.On("GetExerciseMedia", mock.Anything, validID).Return([]models.ExerciseMedia{}, errors.New("database error"))
			},
			expectError: true,
			expected:    nil,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockMongoRepo.ExpectedCalls = nil
			tt.setupMocks()
			
			// Call the method
			media, err := service.GetExerciseMedia(ctx, tt.exerciseID)
			
			// Check result
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, media)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, media)
			}
			
			// Verify mocks
			mockMongoRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteExerciseMedia(t *testing.T) {
	// Create mocks
	mockRepo := new(MockRepository)
	mockMongoRepo := new(MockMongoDBRepository)
	
	// Create service
	service := NewMediaService(mockRepo, mockMongoRepo)
	
	// Create test context with user ID
	ctx := context.WithValue(context.Background(), "user_id", 1)
	
	// Create valid media ID
	validID := primitive.NewObjectID().Hex()
	
	// Setup test cases
	tests := []struct {
		name        string
		mediaID     string
		setupMocks  func()
		expectError bool
	}{
		{
			name:    "Valid request",
			mediaID: validID,
			setupMocks: func() {
				mockMongoRepo.On("DeleteExerciseMedia", mock.Anything, validID).Return(nil)
			},
			expectError: false,
		},
		{
			name:        "Invalid media ID",
			mediaID:     "invalid-id",
			setupMocks:  func() {},
			expectError: true,
		},
		{
			name:    "Database error",
			mediaID: validID,
			setupMocks: func() {
				mockMongoRepo.On("DeleteExerciseMedia", mock.Anything, validID).Return(errors.New("database error"))
			},
			expectError: true,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockMongoRepo.ExpectedCalls = nil
			tt.setupMocks()
			
			// Call the method
			err := service.DeleteExerciseMedia(ctx, tt.mediaID)
			
			// Check result
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			
			// Verify mocks
			mockMongoRepo.AssertExpectations(t)
		})
	}
}