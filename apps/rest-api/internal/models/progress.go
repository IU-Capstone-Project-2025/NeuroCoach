package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutCompletion struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      int                `bson:"user_id" json:"user_id"`
	WorkoutID   primitive.ObjectID `bson:"workout_id" json:"workout_id"`
	CompletedAt time.Time          `bson:"completed_at" json:"completed_at"`
}

type UserProgress struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID            int                `bson:"user_id" json:"user_id"`
	TotalWorkouts     int                `bson:"total_workouts" json:"total_workouts"`
	ConsecutiveDays   int                `bson:"consecutive_days" json:"consecutive_days"`
	Level             string             `bson:"level" json:"level"`
	CompletedWorkouts []string           `bson:"completed_workouts" json:"completed_workouts"`
	LastWorkoutDate   time.Time          `bson:"last_workout_date" json:"last_workout_date"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
}

type CompleteWorkoutRequest struct {
	WorkoutID string `json:"workout_id" validate:"required"`
}
