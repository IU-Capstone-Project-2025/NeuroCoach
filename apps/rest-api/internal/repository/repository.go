package repository

import (
	"context"
	"rest-api/internal/models"
)

type Repository interface {
	// User operations
	CreateUser(ctx context.Context, email, passwordHash string) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// Fitness profile operations
	SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error
	GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error)

	// // Workout plan operations
	// SaveWorkoutPlan(ctx context.Context, userID int, plan *models.WorkoutPlan) error
	// GetWorkoutPlan(ctx context.Context, userID int) (*models.WorkoutPlan, error)

	// // Chat operations
	// SaveChatMessage(ctx context.Context, message *models.ChatMessage) error
	// GetChatHistory(ctx context.Context, userID int, limit int) ([]models.ChatMessage, error)

	// Health check
	Ping(ctx context.Context) error
	Close() error
}

type MongoDBRep interface {
	// Chat operations
	SaveChatMessage(ctx context.Context, message *models.ChatMessage) error
	GetChatHistory(ctx context.Context, userID int) ([]models.ChatMessage, error)

	// Workout plan operations
	SaveWorkoutPlan(ctx context.Context, plan *models.WorkoutPlan) error
	GetWorkoutPlan(ctx context.Context, userID int) (*models.WorkoutPlan, error)
	GetWorkoutByID(ctx context.Context, userID int, workoutID string) (*models.Workout, error)

	// Short plan operations
	SaveShortPlan(ctx context.Context, plan *models.ShortWorkoutPlan) error
	GetShortPlan(ctx context.Context, userID int) (*models.ShortWorkoutPlan, error)

	// Progress operations
	CompleteWorkout(ctx context.Context, userID int, workoutID string) error
	GetUserProgress(ctx context.Context, userID int) (*models.UserProgress, error)
}
