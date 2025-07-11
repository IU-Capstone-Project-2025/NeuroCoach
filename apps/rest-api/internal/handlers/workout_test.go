package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeneratePlan_Handler(t *testing.T) {
	h := &Handlers{AIService: nil}

	req := httptest.NewRequest("POST", "/generate-plan", nil)
	w := httptest.NewRecorder()

	// This should panic due to nil service, so we recover
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to nil service")
		}
	}()

	h.GeneratePlan(w, req)
}

func TestGetWorkoutPlan_Handler(t *testing.T) {
	h := &Handlers{AIService: nil}

	req := httptest.NewRequest("GET", "/workout-plan", nil)
	w := httptest.NewRecorder()

	// This should panic due to nil service, so we recover
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to nil service")
		}
	}()

	h.GetWorkoutPlan(w, req)
}

func TestRegenerateWorkoutPlan_InvalidJSON(t *testing.T) {
	h := &Handlers{}

	req := httptest.NewRequest("POST", "/regenerate-plan", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.RegenerateWorkoutPlan(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetRating_Handler(t *testing.T) {
	h := &Handlers{AIService: nil}

	req := httptest.NewRequest("GET", "/rating", nil)
	w := httptest.NewRecorder()

	// This should panic due to nil service, so we recover
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to nil service")
		}
	}()

	h.GetRating(w, req)
}

func TestGetMotivationalMessage_Handler(t *testing.T) {
	h := &Handlers{AIService: nil}

	req := httptest.NewRequest("GET", "/motivation", nil)
	w := httptest.NewRecorder()

	// This should panic due to nil service, so we recover
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to nil service")
		}
	}()

	h.GetMotivationalMessage(w, req)
}
