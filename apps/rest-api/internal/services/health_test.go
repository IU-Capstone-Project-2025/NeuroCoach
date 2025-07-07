package services

import (
	"context"
	"errors"
	"testing"

	"rest-api/internal/models"
)

// Mock repository for health testing
type mockHealthRepo struct {
	pingError error
}

func (m *mockHealthRepo) CreateUser(ctx context.Context, email, passwordHash string) (int, error) {
	return 1, nil
}

func (m *mockHealthRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (m *mockHealthRepo) SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error {
	return nil
}

func (m *mockHealthRepo) GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error) {
	return nil, nil
}

func (m *mockHealthRepo) Ping(ctx context.Context) error {
	return m.pingError
}

func (m *mockHealthRepo) Close() error {
	return nil
}

func TestHealthService_CheckDB_Success(t *testing.T) {
	repo := &mockHealthRepo{pingError: nil}
	service := NewHealthService(repo)

	err := service.CheckDB(context.Background())

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestHealthService_CheckDB_Error(t *testing.T) {
	expectedError := errors.New("database connection failed")
	repo := &mockHealthRepo{pingError: expectedError}
	service := NewHealthService(repo)

	err := service.CheckDB(context.Background())

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestNewHealthService(t *testing.T) {
	repo := &mockHealthRepo{}
	service := NewHealthService(repo)

	if service == nil {
		t.Error("Expected non-nil service")
	}
}
