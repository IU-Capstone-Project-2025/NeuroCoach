package services

import (
	"context"
	"testing"

	"rest-api/internal/middleware"
)

func TestServiceError_Error(t *testing.T) {
	err := NewServiceError(404, "Not found", nil)

	expected := "Not found: <nil>"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestNewServiceError(t *testing.T) {
	code := 500
	message := "Internal error"

	err := NewServiceError(code, message, nil)

	if err.Code != code {
		t.Errorf("Expected code %d, got %d", code, err.Code)
	}

	if err.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, err.Message)
	}
}

func TestBaseService_GetUserIDFromContext_Success(t *testing.T) {
	service := &BaseService{}

	ctx := context.WithValue(context.Background(), middleware.UserIDKey, 123)

	userID, err := service.GetUserIDFromContext(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if userID != 123 {
		t.Errorf("Expected user ID 123, got %d", userID)
	}
}

func TestBaseService_GetUserIDFromContext_Missing(t *testing.T) {
	service := &BaseService{}

	ctx := context.Background()

	_, err := service.GetUserIDFromContext(ctx)
	if err == nil {
		t.Error("Expected error for missing user ID")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 500 {
			t.Errorf("Expected status code 500, got %d", svcErr.Code)
		}
	}
}
