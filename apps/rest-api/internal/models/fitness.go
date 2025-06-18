package models

import "time"

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
	UserID int `json:"-"`
	Regenerate bool `json:"regenerate"` // Flag to force regeneration
}

type WorkoutPlan struct {
	PlanID    int       `json:"plan_id"`
	UserID    int       `json:"-"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}