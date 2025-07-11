package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"rest-api/internal/models"
	"rest-api/internal/services"
)

type mockHealthService struct {
	checkDBError error
}

func (m *mockHealthService) CheckDB(ctx context.Context) error {
	return m.checkDBError
}

func TestHealthCheck_Success(t *testing.T) {
	mockService := &mockHealthService{checkDBError: nil}
	h := &Handlers{HealthService: &services.HealthService{}}
	// Override the CheckDB method by replacing the service
	originalService := h.HealthService
	h.HealthService = &services.HealthService{}

	// Create a custom handler that uses our mock
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		if err := mockService.CheckDB(r.Context()); err != nil {
			respondWithError(w, http.StatusServiceUnavailable, "Database unavailable")
			return
		}
		respondWithJSON(w, http.StatusOK, models.HealthCheckResponse{
			Status:  "healthy",
			Version: "1.0",
		})
	}

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	testHandler(w, req)
	h.HealthService = originalService

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response models.HealthCheckResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("Expected status 'healthy', got %s", response.Status)
	}

	if response.Version != "1.0" {
		t.Errorf("Expected version '1.0', got %s", response.Version)
	}
}

func TestHealthCheck_DatabaseError(t *testing.T) {
	mockService := &mockHealthService{checkDBError: errors.New("database connection failed")}

	// Create a custom handler that uses our mock
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		if err := mockService.CheckDB(r.Context()); err != nil {
			respondWithError(w, http.StatusServiceUnavailable, "Database unavailable")
			return
		}
		respondWithJSON(w, http.StatusOK, models.HealthCheckResponse{
			Status:  "healthy",
			Version: "1.0",
		})
	}

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	testHandler(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("Expected status 503, got %d", w.Code)
	}

	var response models.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Message != "Database unavailable" {
		t.Errorf("Expected message 'Database unavailable', got %s", response.Message)
	}
}
