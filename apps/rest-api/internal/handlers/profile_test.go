package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveProfile_InvalidJSON(t *testing.T) {
	h := &Handlers{}

	req := httptest.NewRequest("POST", "/profile", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.SaveProfile(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetProfile_Handler(t *testing.T) {
	h := &Handlers{}

	req := httptest.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()

	h.GetProfile(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}
