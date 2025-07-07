package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"rest-api/internal/models"
)

func TestRespondWithJSON_Extended(t *testing.T) {
	w := httptest.NewRecorder()
	data := models.AuthResponse{
		Token: "test-token",
		Email: "test@example.com",
	}

	respondWithJSON(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response models.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Token != data.Token {
		t.Errorf("Expected token %s, got %s", data.Token, response.Token)
	}
}

func TestRespondWithError_Extended(t *testing.T) {
	w := httptest.NewRecorder()

	respondWithError(w, http.StatusNotFound, "Resource not found")

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}

	var response models.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Message != "Resource not found" {
		t.Errorf("Expected message 'Resource not found', got %s", response.Message)
	}

	if response.Error != "Not Found" {
		t.Errorf("Expected error 'Not Found', got %s", response.Error)
	}
}

func TestRegister_InvalidRequestBody(t *testing.T) {
	h := &Handlers{}

	// Test with invalid JSON
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.Register(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestLogin_InvalidRequestBody(t *testing.T) {
	h := &Handlers{}

	// Test with invalid JSON
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.Login(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
