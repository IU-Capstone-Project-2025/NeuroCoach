package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"rest-api/internal/models"
)

type testRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=18"`
}

func TestValidateRequest_Success(t *testing.T) {
	validReq := testRequest{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   25,
	}

	body, _ := json.Marshal(validReq)
	req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	called := false
	handler := ValidateRequest(func(w http.ResponseWriter, r *http.Request, req testRequest) {
		called = true
		if req.Name != validReq.Name {
			t.Errorf("Expected name %s, got %s", validReq.Name, req.Name)
		}
		w.WriteHeader(http.StatusOK)
	})

	handler.ServeHTTP(w, req)

	if !called {
		t.Error("Handler should have been called")
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestValidateRequest_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/test", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler := ValidateRequest(func(w http.ResponseWriter, r *http.Request, req testRequest) {
		t.Error("Handler should not be called")
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response models.ErrorResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	if response.Message != "Invalid request payload" {
		t.Errorf("Expected 'Invalid request payload', got %s", response.Message)
	}
}

func TestValidateRequest_ValidationError(t *testing.T) {
	invalidReq := testRequest{
		Name:  "",
		Email: "invalid-email",
		Age:   15,
	}

	body, _ := json.Marshal(invalidReq)
	req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler := ValidateRequest(func(w http.ResponseWriter, r *http.Request, req testRequest) {
		t.Error("Handler should not be called")
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response models.ErrorResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	if response.Error != "Validation failed" {
		t.Errorf("Expected 'Validation failed', got %s", response.Error)
	}
}
