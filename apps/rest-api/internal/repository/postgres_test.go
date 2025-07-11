package repository

import (
	"context"
	"testing"
	"time"

	"rest-api/internal/models"
)

// Mock PostgreSQL repository for testing
type mockPostgresRepo struct {
	users    map[string]*models.User
	profiles map[int]*models.FitnessProfile
	nextID   int
	pingErr  error
}

func newMockPostgresRepo() *mockPostgresRepo {
	return &mockPostgresRepo{
		users:    make(map[string]*models.User),
		profiles: make(map[int]*models.FitnessProfile),
		nextID:   1,
	}
}

func (m *mockPostgresRepo) CreateUser(ctx context.Context, email, passwordHash string) (int, error) {
	if _, exists := m.users[email]; exists {
		return 0, ErrNotFound
	}

	user := &models.User{
		ID:           m.nextID,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
	m.users[email] = user
	m.nextID++
	return user.ID, nil
}

func (m *mockPostgresRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if user, exists := m.users[email]; exists {
		return user, nil
	}
	return nil, ErrNotFound
}

func (m *mockPostgresRepo) SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error {
	m.profiles[userID] = profile
	return nil
}

func (m *mockPostgresRepo) GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error) {
	if profile, exists := m.profiles[userID]; exists {
		return profile, nil
	}
	return nil, ErrNotFound
}

func (m *mockPostgresRepo) Ping(ctx context.Context) error {
	return m.pingErr
}

func (m *mockPostgresRepo) Close() error {
	return nil
}

func TestMockPostgresRepo_CreateUser(t *testing.T) {
	repo := newMockPostgresRepo()

	userID, err := repo.CreateUser(context.Background(), "test@example.com", "hashedpassword")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if userID != 1 {
		t.Errorf("Expected user ID 1, got %d", userID)
	}

	// Test duplicate email
	_, err = repo.CreateUser(context.Background(), "test@example.com", "hashedpassword")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound for duplicate email, got %v", err)
	}
}

func TestMockPostgresRepo_GetUserByEmail(t *testing.T) {
	repo := newMockPostgresRepo()

	// Create user first
	_, _ = repo.CreateUser(context.Background(), "test@example.com", "hashedpassword")

	user, err := repo.GetUserByEmail(context.Background(), "test@example.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", user.Email)
	}

	// Test non-existent user
	_, err = repo.GetUserByEmail(context.Background(), "nonexistent@example.com")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}

func TestMockPostgresRepo_FitnessProfile(t *testing.T) {
	repo := newMockPostgresRepo()

	profile := &models.FitnessProfile{
		Height:           175.0,
		Weight:           70.0,
		Age:              25,
		Goal:             "weight_loss",
		Timeframe:        "3months",
		FitnessLevel:     "intermediate",
		AvailableMinutes: 180,
		UpdatedAt:        time.Now(),
	}

	// Save profile
	err := repo.SaveFitnessProfile(context.Background(), 1, profile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Get profile
	savedProfile, err := repo.GetFitnessProfile(context.Background(), 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if savedProfile.Height != profile.Height {
		t.Errorf("Expected height %f, got %f", profile.Height, savedProfile.Height)
	}

	// Test non-existent profile
	_, err = repo.GetFitnessProfile(context.Background(), 999)
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}
