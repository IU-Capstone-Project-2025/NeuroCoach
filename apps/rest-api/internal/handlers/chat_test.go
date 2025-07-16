package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChat_InvalidJSON(t *testing.T) {
	h := &Handlers{}

	req := httptest.NewRequest("POST", "/chat", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.Chat(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetChatHistory_Handler(t *testing.T) {
	h := &Handlers{AIService: nil}

	req := httptest.NewRequest("GET", "/chat/history", nil)
	w := httptest.NewRecorder()

	// This should panic due to nil service, so we recover
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to nil service")
		}
	}()

	h.GetChatHistory(w, req)
}

func TestCompleteWorkout_InvalidJSON(t *testing.T) {
	h := &Handlers{}

	req := httptest.NewRequest("POST", "/complete-workout", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.CompleteWorkout(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetUserProgress_Handler(t *testing.T) {
	h := &Handlers{AIService: nil}

	req := httptest.NewRequest("GET", "/progress", nil)
	w := httptest.NewRecorder()

	// This should panic due to nil service, so we recover
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to nil service")
		}
	}()

	h.GetUserProgress(w, req)
}
