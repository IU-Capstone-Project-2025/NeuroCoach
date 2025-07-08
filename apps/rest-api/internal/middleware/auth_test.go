package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"rest-api/pkg/utils"
)

func TestAuthMiddleware_Success(t *testing.T) {
	secret := "test-secret"
	userID := 123
	email := "test@example.com"

	// Generate a valid token
	token, err := utils.GenerateJWT(userID, email, secret, time.Hour)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Create a test handler that checks if user ID is in context
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		extractedUserID, ok := GetUserIDFromContext(r.Context())
		if !ok {
			t.Error("Expected user ID in context")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if extractedUserID != userID {
			t.Errorf("Expected user ID %d, got %d", userID, extractedUserID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success"))
	})

	// Wrap with auth middleware
	authHandler := AuthMiddleware(secret)(testHandler)

	// Create request with Authorization header
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Execute request
	authHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "success" {
		t.Errorf("Expected 'success', got %s", w.Body.String())
	}
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	secret := "test-secret"

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called")
	})

	authHandler := AuthMiddleware(secret)(testHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	authHandler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidTokenFormat(t *testing.T) {
	secret := "test-secret"

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called")
	})

	authHandler := AuthMiddleware(secret)(testHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "InvalidToken")
	w := httptest.NewRecorder()

	authHandler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	secret := "test-secret"

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called")
	})

	authHandler := AuthMiddleware(secret)(testHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")
	w := httptest.NewRecorder()

	authHandler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_ExpiredToken(t *testing.T) {
	secret := "test-secret"
	userID := 123
	email := "test@example.com"

	// Generate an expired token
	token, err := utils.GenerateJWT(userID, email, secret, -time.Hour)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called")
	})

	authHandler := AuthMiddleware(secret)(testHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	authHandler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestGetUserIDFromContext_Success(t *testing.T) {
	userID := 123
	ctx := context.WithValue(context.Background(), UserIDKey, userID)

	extractedUserID, ok := GetUserIDFromContext(ctx)

	if !ok {
		t.Error("Expected to find user ID in context")
	}

	if extractedUserID != userID {
		t.Errorf("Expected user ID %d, got %d", userID, extractedUserID)
	}
}

func TestGetUserIDFromContext_NotFound(t *testing.T) {
	ctx := context.Background()

	_, ok := GetUserIDFromContext(ctx)

	if ok {
		t.Error("Expected not to find user ID in context")
	}
}

func TestGetUserIDFromContext_WrongType(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDKey, "not-an-int")

	_, ok := GetUserIDFromContext(ctx)

	if ok {
		t.Error("Expected not to find user ID with wrong type")
	}
}
