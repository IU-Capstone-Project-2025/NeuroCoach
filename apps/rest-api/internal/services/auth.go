package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"rest-api/internal/models"
	"rest-api/internal/repository"
	"rest-api/pkg/utils"
)

type AuthService struct {
	BaseService
	JWTSecret string
	JWTExpiry time.Duration
}

func NewAuthService(repo repository.Repository, jwtSecret string, jwtExpiry time.Duration) *AuthService {
	return &AuthService{
		BaseService: BaseService{Repo: repo},
		JWTSecret:   jwtSecret,
		JWTExpiry:   jwtExpiry,
	}
}

func (s *AuthService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	existing, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to check existing user",
			err,
		)
	}
	if existing != nil {
		return nil, NewServiceError(
			http.StatusConflict,
			"Email already registered",
			nil,
		)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to hash password",
			err,
		)
	}

	// Create user
	userID, err := s.Repo.CreateUser(ctx, req.Email, hashedPassword)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to create user",
			err,
		)
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(userID, req.Email, s.JWTSecret, s.JWTExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate token",
			err,
		)
	}

	return &models.AuthResponse{
		Token: token,
		Email: req.Email,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	// Get user by email
	user, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, NewServiceError(
				http.StatusUnauthorized,
				"Invalid credentials",
				nil,
			)
		}
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to retrieve user",
			err,
		)
	}

	// Compare password
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, NewServiceError(
			http.StatusUnauthorized,
			"Invalid credentials",
			nil,
		)
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, s.JWTSecret, s.JWTExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate token",
			err,
		)
	}

	return &models.AuthResponse{
		Token: token,
		Email: user.Email,
	}, nil
}
