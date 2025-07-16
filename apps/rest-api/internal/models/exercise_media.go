package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseMedia struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	ImageURL    string             `bson:"image_url" json:"image_url"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type ExerciseMediaRequest struct {
	Name        string `json:"name" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required,url"`
	Description string `json:"description" validate:"required"`
}