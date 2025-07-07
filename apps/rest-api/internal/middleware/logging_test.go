package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoggingMiddleware(t *testing.T) {
	// Create a test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test response"))
	})

	// Wrap with logging middleware
	loggedHandler := LoggingMiddleware(testHandler)

	// Create test request
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Execute request
	loggedHandler.ServeHTTP(w, req)

	// Check that the original handler was called
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "test response" {
		t.Errorf("Expected 'test response', got %s", w.Body.String())
	}
}

func TestLoggingMiddleware_CustomStatus(t *testing.T) {
	// Create a test handler that returns a custom status
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	})

	// Wrap with logging middleware
	loggedHandler := LoggingMiddleware(testHandler)

	// Create test request
	req := httptest.NewRequest("POST", "/api/test", nil)
	w := httptest.NewRecorder()

	// Execute request
	loggedHandler.ServeHTTP(w, req)

	// Check that the custom status was preserved
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}

	if w.Body.String() != "not found" {
		t.Errorf("Expected 'not found', got %s", w.Body.String())
	}
}

func TestResponseWriter_WriteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	rw := &responseWriter{w, http.StatusOK}

	// Test that WriteHeader sets the status
	rw.WriteHeader(http.StatusCreated)

	if rw.status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rw.status)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Expected response code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestLoggingMiddleware_Timing(t *testing.T) {
	// Create a test handler that takes some time
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	})

	// Wrap with logging middleware
	loggedHandler := LoggingMiddleware(testHandler)

	// Create test request
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Measure execution time
	start := time.Now()
	loggedHandler.ServeHTTP(w, req)
	duration := time.Since(start)

	// Check that the handler was called and took some time
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if duration < 10*time.Millisecond {
		t.Errorf("Expected duration >= 10ms, got %v", duration)
	}
}
