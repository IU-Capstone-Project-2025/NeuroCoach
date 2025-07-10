package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseMedia struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ExerciseID  primitive.ObjectID `bson:"exercise_id" json:"exercise_id"`
	ImageURL    string             `bson:"image_url" json:"image_url"`
	Description string             `bson:"description" json:"description"`
	Order       int                `bson:"order" json:"order"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type ExerciseMediaRequest struct {
	ExerciseID  string `json:"exercise_id" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required,url"`
	Description string `json:"description" validate:"required"`
	Order       int    `json:"order" validate:"required,gte=0"`
}