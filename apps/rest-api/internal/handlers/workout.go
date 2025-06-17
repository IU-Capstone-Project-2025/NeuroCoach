package handlers

import (
	"net/http"
)

func (h *Handlers) GeneratePlan(w http.ResponseWriter, r *http.Request) {
	plan, err := h.AIService.GenerateWorkoutPlan(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"plan": plan})
}

func (h *Handlers) GetWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	plan, err := h.AIService.GenerateWorkoutPlan(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, plan)
}