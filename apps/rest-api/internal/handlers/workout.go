package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api/internal/models"
)

func (h *Handlers) GeneratePlan(w http.ResponseWriter, r *http.Request) {
	plan, err := h.AIService.GenerateWorkoutPlan(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, plan)
}

func (h *Handlers) GetWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	plan, err := h.AIService.GenerateWorkoutPlan(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, plan)
}

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
