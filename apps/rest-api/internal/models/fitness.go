package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FitnessProfile struct {
	UserID           int       `json:"-"`
	Height           float64   `json:"height" validate:"required,gt=0"`
	Weight           float64   `json:"weight" validate:"required,gt=0"`
	Age              int       `json:"age" validate:"required,gte=13,lte=120"`
	Goal             string    `json:"goal" validate:"required,oneof=weight_loss muscle_gain endurance flexibility general_fitness"`
	HealthIssues     []string  `json:"health_issues"`
	Timeframe        string    `json:"timeframe" validate:"required,oneof=1month 3months 6months 1year"`
	FitnessLevel     string    `json:"fitness_level" validate:"required,oneof=beginner intermediate advanced"`
	AvailableMinutes int       `json:"available_minutes" validate:"required,gte=30,lte=1000"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type WorkoutPlanRequest struct {
	UserID     int  `json:"-"`
	Regenerate bool `json:"regenerate"` // Flag to force regeneration
}

type RegenerateWorkoutPlanRequest struct {
	Comments string `json:"comments" validate:"required,max=1000"`
}

type WorkoutPlan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    int                `bson:"user_id" json:"user_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Status    bool               `bson:"status" json:"status"`
	Title     string             `bson:"title" json:"title"`
	Workouts  []Workout          `bson:"workouts" json:"workouts"`
}

type ShortWorkoutPlan struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID          int                `bson:"user_id" json:"user_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
	Status          bool               `bson:"status" json:"status"`
	Title           string             `bson:"title" json:"title"`
	BaseWorkouts    []Workout          `bson:"base_workouts" json:"base_workouts"`
	Timeframe       string             `bson:"timeframe" json:"timeframe"`
	WorkoutsPerWeek int                `bson:"workouts_per_week" json:"workouts_per_week"`
}

type Workout struct {
	WorkoutID     primitive.ObjectID `bson:"workout_id,omitempty" json:"workout_id"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description,omitempty" json:"description,omitempty"`
	Status        string             `bson:"status" json:"status"`
	ScheduledDate time.Time          `bson:"scheduled_date" json:"scheduled_date"`
	Exercises     []Exercise         `bson:"exercises" json:"exercises"`
}

type Exercise struct {
	ExerciseID  primitive.ObjectID `bson:"exercise_id,omitempty" json:"exercise_id"`
	Name        string             `bson:"name" json:"name"`
	MuscleGroup string             `bson:"muscle_group" json:"muscle_group"`
	Sets        int                `bson:"sets" json:"sets"`
	Reps        int                `bson:"reps" json:"reps"`
	RestSec     int                `bson:"rest_sec,omitempty" json:"rest_sec,omitempty"`
	Notes       string             `bson:"notes,omitempty" json:"notes,omitempty"`
	Technique   string             `bson:"technique,omitempty" json:"technique,omitempty"`
}
