package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api/internal/models"
)

// Chat godoc
// @Summary Chat with AI
// @Description Send message to AI assistant
// @Tags chat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.ChatRequest true "Chat message"
// @Success 200 {object} models.ChatResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/chat [post]
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

// GetChatHistory godoc
// @Summary Get chat history
// @Description Get user's chat history with AI
// @Tags chat
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ChatHistory
// @Failure 401 {object} models.ErrorResponse
// @Router /api/chat/history [get]
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

// CompleteWorkout godoc
// @Summary Complete workout
// @Description Mark a workout as completed
// @Tags workout
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CompleteWorkoutRequest true "Workout completion data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/complete-workout [post]
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

// GetUserProgress godoc
// @Summary Get user progress
// @Description Get user's workout progress and statistics
// @Tags progress
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.UserProgress
// @Failure 401 {object} models.ErrorResponse
// @Router /api/progress [get]
func (h *Handlers) GetUserProgress(w http.ResponseWriter, r *http.Request) {
	progress, err := h.AIService.GetUserProgress(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, progress)
}
