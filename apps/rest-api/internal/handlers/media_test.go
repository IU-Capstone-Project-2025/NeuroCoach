package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"rest-api/internal/models"
	"rest-api/internal/services"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock MediaService
type MockMediaService struct {
	mock.Mock
}

func (m *MockMediaService) SaveExerciseMedia(ctx interface{}, req *models.ExerciseMediaRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockMediaService) GetExerciseMedia(ctx interface{}, exerciseID string) ([]models.ExerciseMedia, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).([]models.ExerciseMedia), args.Error(1)
}

func (m *MockMediaService) DeleteExerciseMedia(ctx interface{}, mediaID string) error {
	args := m.Called(ctx, mediaID)
	return args.Error(0)
}

func TestSaveExerciseMediaHandler(t *testing.T) {
	// Create mock service
	mockService := new(MockMediaService)
	
	// Create handlers with mock service
	h := &Handlers{
		MediaService: mockService,
	}
	
	// Create test cases
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func()
		expectedStatus int
	}{
		{
			name: "Valid request",
			requestBody: models.ExerciseMediaRequest{
				ExerciseID:  "60f7b5b9e6b3f3b3e8b4b5b9",
				ImageURL:    "https://example.com/image.jpg",
				Description: "Test description",
				Order:       1,
			},
			setupMocks: func() {
				mockService.On("SaveExerciseMedia", mock.Anything, mock.AnythingOfType("*models.ExerciseMediaRequest")).Return(nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Invalid request body",
			requestBody: struct {
				Invalid string `json:"invalid"`
			}{
				Invalid: "invalid",
			},
			setupMocks:     func() {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Service error",
			requestBody: models.ExerciseMediaRequest{
				ExerciseID:  "60f7b5b9e6b3f3b3e8b4b5b9",
				ImageURL:    "https://example.com/image.jpg",
				Description: "Test description",
				Order:       1,
			},
			setupMocks: func() {
				mockService.On("SaveExerciseMedia", mock.Anything, mock.AnythingOfType("*models.ExerciseMediaRequest")).Return(
					services.NewServiceError(http.StatusInternalServerError, "Database error", errors.New("db error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockService.ExpectedCalls = nil
			tt.setupMocks()
			
			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/api/exercise/media", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			
			// Create response recorder
			rr := httptest.NewRecorder()
			
			// Call handler
			h.SaveExerciseMedia(rr, req)
			
			// Check response
			assert.Equal(t, tt.expectedStatus, rr.Code)
			
			// Verify mocks
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetExerciseMediaHandler(t *testing.T) {
	// Create mock service
	mockService := new(MockMediaService)
	
	// Create handlers with mock service
	h := &Handlers{
		MediaService: mockService,
	}
	
	// Create sample media
	sampleMedia := []models.ExerciseMedia{}
	
	// Create test cases
	tests := []struct {
		name           string
		exerciseID     string
		setupMocks     func()
		expectedStatus int
	}{
		{
			name:       "Valid request",
			exerciseID: "60f7b5b9e6b3f3b3e8b4b5b9",
			setupMocks: func() {
				mockService.On("GetExerciseMedia", mock.Anything, "60f7b5b9e6b3f3b3e8b4b5b9").Return(sampleMedia, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "Service error",
			exerciseID: "60f7b5b9e6b3f3b3e8b4b5b9",
			setupMocks: func() {
				mockService.On("GetExerciseMedia", mock.Anything, "60f7b5b9e6b3f3b3e8b4b5b9").Return(
					[]models.ExerciseMedia{},
					services.NewServiceError(http.StatusInternalServerError, "Database error", errors.New("db error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockService.ExpectedCalls = nil
			tt.setupMocks()
			
			// Create request
			req, _ := http.NewRequest("GET", "/api/exercise/"+tt.exerciseID+"/media", nil)
			
			// Create response recorder
			rr := httptest.NewRecorder()
			
			// Create router with vars
			router := mux.NewRouter()
			router.HandleFunc("/api/exercise/{exercise_id}/media", h.GetExerciseMedia)
			
			// Call handler
			router.ServeHTTP(rr, req)
			
			// Check response
			assert.Equal(t, tt.expectedStatus, rr.Code)
			
			// Verify mocks
			mockService.AssertExpectations(t)
		})
	}
}

func TestDeleteExerciseMediaHandler(t *testing.T) {
	// Create mock service
	mockService := new(MockMediaService)
	
	// Create handlers with mock service
	h := &Handlers{
		MediaService: mockService,
	}
	
	// Create test cases
	tests := []struct {
		name           string
		mediaID        string
		setupMocks     func()
		expectedStatus int
	}{
		{
			name:    "Valid request",
			mediaID: "60f7b5b9e6b3f3b3e8b4b5b9",
			setupMocks: func() {
				mockService.On("DeleteExerciseMedia", mock.Anything, "60f7b5b9e6b3f3b3e8b4b5b9").Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:    "Service error",
			mediaID: "60f7b5b9e6b3f3b3e8b4b5b9",
			setupMocks: func() {
				mockService.On("DeleteExerciseMedia", mock.Anything, "60f7b5b9e6b3f3b3e8b4b5b9").Return(
					services.NewServiceError(http.StatusInternalServerError, "Database error", errors.New("db error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockService.ExpectedCalls = nil
			tt.setupMocks()
			
			// Create request
			req, _ := http.NewRequest("DELETE", "/api/exercise/media/"+tt.mediaID, nil)
			
			// Create response recorder
			rr := httptest.NewRecorder()
			
			// Create router with vars
			router := mux.NewRouter()
			router.HandleFunc("/api/exercise/media/{media_id}", h.DeleteExerciseMedia)
			
			// Call handler
			router.ServeHTTP(rr, req)
			
			// Check response
			assert.Equal(t, tt.expectedStatus, rr.Code)
			
			// Verify mocks
			mockService.AssertExpectations(t)
		})
	}
}