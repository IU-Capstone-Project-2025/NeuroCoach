package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    int                `bson:"user_id" json:"user_id"`
	Message   string             `bson:"message" json:"message"`
	Response  string             `bson:"response" json:"response"`
	IsUser    bool               `bson:"is_user" json:"is_user"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type ChatRequest struct {
	Message string `json:"message" validate:"required,max=500"`
}

type ChatHistory struct {
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Response string `json:"response"`
}
