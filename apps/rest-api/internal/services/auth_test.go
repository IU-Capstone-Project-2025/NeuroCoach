package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"rest-api/internal/models"
	"rest-api/internal/repository"
)

// Mock repository for auth testing
type mockAuthRepo struct {
	users  map[string]*models.User
	nextID int
}

func newMockAuthRepo() *mockAuthRepo {
	return &mockAuthRepo{
		users:  make(map[string]*models.User),
		nextID: 1,
	}
}

func (m *mockAuthRepo) CreateUser(ctx context.Context, email, passwordHash string) (int, error) {
	if _, exists := m.users[email]; exists {
		return 0, errors.New("user already exists")
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

func (m *mockAuthRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if user, exists := m.users[email]; exists {
		return user, nil
	}
	return nil, repository.ErrNotFound
}

func (m *mockAuthRepo) SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error {
	return nil
}

func (m *mockAuthRepo) GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error) {
	return nil, repository.ErrNotFound
}

func (m *mockAuthRepo) Ping(ctx context.Context) error {
	return nil
}

func (m *mockAuthRepo) Close() error {
	return nil
}

func TestAuthService_Register_Success(t *testing.T) {
	repo := newMockAuthRepo()
	service := NewAuthService(repo, "test-secret", time.Hour)

	req := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := service.Register(context.Background(), req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	if resp.Email != req.Email {
		t.Errorf("Expected email %s, got %s", req.Email, resp.Email)
	}

	if resp.Token == "" {
		t.Error("Expected non-empty token")
	}
}

func TestAuthService_Register_DuplicateEmail(t *testing.T) {
	repo := newMockAuthRepo()
	service := NewAuthService(repo, "test-secret", time.Hour)

	req := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// First registration should succeed
	_, err := service.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("First registration failed: %v", err)
	}

	// Second registration should fail
	_, err = service.Register(context.Background(), req)
	if err == nil {
		t.Error("Expected error for duplicate email")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 409 {
			t.Errorf("Expected status code 409, got %d", svcErr.Code)
		}
	}
}

func TestAuthService_Login_Success(t *testing.T) {
	repo := newMockAuthRepo()
	service := NewAuthService(repo, "test-secret", time.Hour)

	// Register user first
	registerReq := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	_, err := service.Register(context.Background(), registerReq)
	if err != nil {
		t.Fatalf("Registration failed: %v", err)
	}

	// Now login
	loginReq := models.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := service.Login(context.Background(), loginReq)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	if resp.Email != loginReq.Email {
		t.Errorf("Expected email %s, got %s", loginReq.Email, resp.Email)
	}

	if resp.Token == "" {
		t.Error("Expected non-empty token")
	}
}

func TestAuthService_Login_InvalidCredentials(t *testing.T) {
	repo := newMockAuthRepo()
	service := NewAuthService(repo, "test-secret", time.Hour)

	// Register user first
	registerReq := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	_, err := service.Register(context.Background(), registerReq)
	if err != nil {
		t.Fatalf("Registration failed: %v", err)
	}

	// Try login with wrong password
	loginReq := models.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	_, err = service.Login(context.Background(), loginReq)
	if err == nil {
		t.Error("Expected error for invalid credentials")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 401 {
			t.Errorf("Expected status code 401, got %d", svcErr.Code)
		}
	}
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	repo := newMockAuthRepo()
	service := NewAuthService(repo, "test-secret", time.Hour)

	loginReq := models.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	_, err := service.Login(context.Background(), loginReq)
	if err == nil {
		t.Error("Expected error for non-existent user")
	}

	if svcErr, ok := err.(ServiceError); ok {
		if svcErr.Code != 401 {
			t.Errorf("Expected status code 401, got %d", svcErr.Code)
		}
	}
}
