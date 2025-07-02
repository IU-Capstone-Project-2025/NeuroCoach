package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"rest-api/internal/models"
)

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"message": "test"}

	respondWithJSON(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}

func TestRespondWithError(t *testing.T) {
	w := httptest.NewRecorder()

	respondWithError(w, http.StatusBadRequest, "test error")

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response models.ErrorResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	if response.Message != "test error" {
		t.Errorf("Expected message 'test error', got %s", response.Message)
	}
}

func TestRegister_RequestStructure(t *testing.T) {
	reqBody := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))

	// Простой тест структуры запроса
	if req.Method != "POST" {
		t.Errorf("Expected POST method, got %s", req.Method)
	}

	if req.URL.Path != "/register" {
		t.Errorf("Expected /register path, got %s", req.URL.Path)
	}

	// Проверяем, что можем декодировать тело запроса
	var decodedBody models.RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&decodedBody)
	if err != nil {
		t.Errorf("Failed to decode request body: %v", err)
	}

	if decodedBody.Email != reqBody.Email {
		t.Errorf("Expected email %s, got %s", reqBody.Email, decodedBody.Email)
	}
}

func TestNewHandlers(t *testing.T) {
	h := NewHandlers(nil, nil, nil, nil)

	if h == nil {
		t.Error("Expected non-nil handlers")
	}
}
