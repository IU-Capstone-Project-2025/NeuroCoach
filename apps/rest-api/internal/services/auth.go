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
	JWTSecret     string
	JWTExpiry     time.Duration
	RefreshExpiry time.Duration
}

func NewAuthService(repo repository.Repository, jwtSecret string, jwtExpiry, refreshExpiry time.Duration) *AuthService {
	return &AuthService{
		BaseService:   BaseService{Repo: repo},
		JWTSecret:     jwtSecret,
		JWTExpiry:     jwtExpiry,
		RefreshExpiry: refreshExpiry,
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

	// Generate access and refresh tokens
	accessToken, err := utils.GenerateJWT(userID, req.Email, s.JWTSecret, s.JWTExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate access token",
			err,
		)
	}

	refreshToken, err := utils.GenerateRefreshToken(userID, s.JWTSecret, s.RefreshExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate refresh token",
			err,
		)
	}

	return &models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.JWTExpiry.Seconds()),
		Email:        req.Email,
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

	// Generate access and refresh tokens
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, s.JWTSecret, s.JWTExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate access token",
			err,
		)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, s.JWTSecret, s.RefreshExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate refresh token",
			err,
		)
	}

	return &models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.JWTExpiry.Seconds()),
		Email:        user.Email,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req models.RefreshTokenRequest) (*models.AuthResponse, error) {
	// Validate refresh token
	claims, err := utils.ValidateJWT(req.RefreshToken, s.JWTSecret)
	if err != nil {
		return nil, NewServiceError(
			http.StatusUnauthorized,
			"Invalid refresh token",
			err,
		)
	}

	// Check if it's a refresh token
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
		return nil, NewServiceError(
			http.StatusUnauthorized,
			"Invalid token type",
			nil,
		)
	}

	// Extract user ID
	userID, err := utils.GetUserIDFromClaims(claims)
	if err != nil {
		return nil, NewServiceError(
			http.StatusUnauthorized,
			"Invalid user ID in token",
			err,
		)
	}

	// Get user details
	user, err := s.Repo.GetUserByEmail(ctx, "") // We'll need to modify this
	if err != nil {
		return nil, NewServiceError(
			http.StatusUnauthorized,
			"User not found",
			err,
		)
	}

	// Generate new tokens
	accessToken, err := utils.GenerateJWT(userID, user.Email, s.JWTSecret, s.JWTExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate access token",
			err,
		)
	}

	newRefreshToken, err := utils.GenerateRefreshToken(userID, s.JWTSecret, s.RefreshExpiry)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to generate refresh token",
			err,
		)
	}

	return &models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.JWTExpiry.Seconds()),
		Email:        user.Email,
	}, nil
}
