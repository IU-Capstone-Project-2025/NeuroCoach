package handlers

import (
	"net/http"
	"encoding/json"

	"rest-api/internal/models"
)

func (h *Handlers) Chat(w http.ResponseWriter, r *http.Request) {
	var req models.ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	response, err := h.AIService.Chat(r.Context(), req.Message)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, models.ChatResponse{
		Response: response,
	})
}

func (h *Handlers) GetChatHistory(w http.ResponseWriter, r *http.Request) {
	history, err := h.AIService.GetChatHistory(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, models.ChatHistory{
		Messages: history,
	})
}

func (h *Handlers) CompleteWorkout(w http.ResponseWriter, r *http.Request) {
	var req models.CompleteWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	if err := h.AIService.CompleteWorkout(r.Context(), req.WorkoutID); err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Workout completed successfully"})
}

func (h *Handlers) GetUserProgress(w http.ResponseWriter, r *http.Request) {
	progress, err := h.AIService.GetUserProgress(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, progress)
}