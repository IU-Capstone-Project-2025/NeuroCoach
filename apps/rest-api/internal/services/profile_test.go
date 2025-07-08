package services

import (
	"context"
	"testing"
	"time"

	"rest-api/internal/middleware"
	"rest-api/internal/models"
	"rest-api/internal/repository"
)

// Mock repository for profile testing
type mockProfileRepo struct {
	profiles map[int]*models.FitnessProfile
}

func newMockProfileRepo() *mockProfileRepo {
	return &mockProfileRepo{
		profiles: make(map[int]*models.FitnessProfile),
	}
}

func (m *mockProfileRepo) CreateUser(ctx context.Context, email, passwordHash string) (int, error) {
	return 1, nil
}

func (m *mockProfileRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{ID: 1, Email: email}, nil
}

func (m *mockProfileRepo) SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error {
	m.profiles[userID] = profile
	return nil
}

func (m *mockProfileRepo) GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error) {
	if profile, exists := m.profiles[userID]; exists {
		return profile, nil
	}
	return nil, repository.ErrNotFound
}

func (m *mockProfileRepo) Ping(ctx context.Context) error {
	return nil
}

func (m *mockProfileRepo) Close() error {
	return nil
}

func TestProfileService_SaveProfile_Success(t *testing.T) {
	repo := newMockProfileRepo()
	service := NewProfileService(repo)

	profile := models.FitnessProfile{
		Height:           175.0,
		Weight:           70.0,
		Age:              25,
		Goal:             "weight_loss",
		HealthIssues:     []string{"knee_pain"},
		Timeframe:        "3months",
		FitnessLevel:     "intermediate",
		AvailableMinutes: 180,
		UpdatedAt:        time.Now(),
	}

	// Create context with user ID
	ctx := context.WithValue(context.Background(), middleware.UserIDKey, 1)

	err := service.SaveProfile(ctx, profile)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify profile was saved
	savedProfile, exists := repo.profiles[1]
	if !exists {
		t.Error("Profile was not saved")
	}

	if savedProfile.Height != profile.Height {
		t.Errorf("Expected height %f, got %f", profile.Height, savedProfile.Height)
	}

	if savedProfile.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", savedProfile.UserID)
	}
}

func TestProfileService_SaveProfile_NoUserID(t *testing.T) {
	repo := newMockProfileRepo()
	service := NewProfileService(repo)

	profile := models.FitnessProfile{
		Height:           175.0,
		Weight:           70.0,
		Age:              25,
		Goal:             "weight_loss",
		Timeframe:        "3months",
		FitnessLevel:     "intermediate",
		AvailableMinutes: 180,
	}

	// Context without user ID
	ctx := context.Background()

	err := service.SaveProfile(ctx, profile)

	if err == nil {
		t.Error("Expected error for missing user ID")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 500 {
			t.Errorf("Expected status code 500, got %d", svcErr.Code)
		}
	}
}

func TestProfileService_GetProfile_Success(t *testing.T) {
	repo := newMockProfileRepo()
	service := NewProfileService(repo)

	// Save a profile first
	originalProfile := &models.FitnessProfile{
		Height:           175.0,
		Weight:           70.0,
		Age:              25,
		Goal:             "weight_loss",
		HealthIssues:     []string{"knee_pain"},
		Timeframe:        "3months",
		FitnessLevel:     "intermediate",
		AvailableMinutes: 180,
		UpdatedAt:        time.Now(),
	}
	repo.profiles[1] = originalProfile

	// Create context with user ID
	ctx := context.WithValue(context.Background(), middleware.UserIDKey, 1)

	profile, err := service.GetProfile(ctx)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if profile == nil {
		t.Fatal("Expected profile, got nil")
	}

	if profile.Height != originalProfile.Height {
		t.Errorf("Expected height %f, got %f", originalProfile.Height, profile.Height)
	}

	if profile.Goal != originalProfile.Goal {
		t.Errorf("Expected goal %s, got %s", originalProfile.Goal, profile.Goal)
	}
}

func TestProfileService_GetProfile_NotFound(t *testing.T) {
	repo := newMockProfileRepo()
	service := NewProfileService(repo)

	// Create context with user ID but no profile saved
	ctx := context.WithValue(context.Background(), middleware.UserIDKey, 1)

	_, err := service.GetProfile(ctx)

	if err == nil {
		t.Error("Expected error for profile not found")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 404 {
			t.Errorf("Expected status code 404, got %d", svcErr.Code)
		}
	}
}

func TestProfileService_GetProfile_NoUserID(t *testing.T) {
	repo := newMockProfileRepo()
	service := NewProfileService(repo)

	// Context without user ID
	ctx := context.Background()

	_, err := service.GetProfile(ctx)

	if err == nil {
		t.Error("Expected error for missing user ID")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 500 {
			t.Errorf("Expected status code 500, got %d", svcErr.Code)
		}
	}
}
