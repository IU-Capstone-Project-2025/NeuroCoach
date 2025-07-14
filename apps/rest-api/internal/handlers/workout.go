package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api/internal/models"
)

// GeneratePlan godoc
// @Summary Generate workout plan
// @Description Generate a new workout plan based on user profile
// @Tags workout
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.WorkoutPlan
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/generate-plan [post]
func (h *Handlers) GeneratePlan(w http.ResponseWriter, r *http.Request) {
	plan, err := h.AIService.GenerateWorkoutPlan(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, plan)
}

// GetWorkoutPlan godoc
// @Summary Get workout plan
// @Description Get user's current workout plan
// @Tags workout
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.WorkoutPlan
// @Failure 401 {object} models.ErrorResponse
// @Router /api/workout-plan [get]
func (h *Handlers) GetWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	plan, err := h.AIService.GenerateWorkoutPlan(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, plan)
}

// RegenerateWorkoutPlan godoc
// @Summary Regenerate workout plan
// @Description Regenerate workout plan based on user feedback
// @Tags workout
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.RegenerateWorkoutPlanRequest true "Regeneration feedback"
// @Success 200 {object} models.WorkoutPlan
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/regenerate-plan [post]
func (h *Handlers) RegenerateWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	var req models.RegenerateWorkoutPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	plan, err := h.AIService.RegenerateWorkoutPlan(r.Context(), req.Comments)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, plan)
}

// GetRating godoc
// @Summary Get user rating
// @Description Get user rating and leaderboard
// @Tags rating
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.UserRating
// @Failure 401 {object} models.ErrorResponse
// @Router /api/rating [get]
func (h *Handlers) GetRating(w http.ResponseWriter, r *http.Request) {
	rating, err := h.AIService.GetRating(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, rating)
}

// GetMotivationalMessage godoc
// @Summary Get motivational message
// @Description Get AI-generated motivational message
// @Tags motivation
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} models.ErrorResponse
// @Router /api/motivation [get]
func (h *Handlers) GetMotivationalMessage(w http.ResponseWriter, r *http.Request) {
	message, err := h.AIService.GenerateMotivationalMessage(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": message})
}
