package models

import "time"

type ChatMessage struct {
	ID        int       `json:"id"`
	UserID    int       `json:"-"`
	Message   string    `json:"message" validate:"required,max=500"`
	Response  string    `json:"response"`
	IsUser    bool      `json:"is_user"`
	CreatedAt time.Time `json:"created_at"`
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