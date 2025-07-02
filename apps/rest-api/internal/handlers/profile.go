package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api/internal/middleware"
	"rest-api/internal/models"
)

func (h *Handlers) SaveProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.FitnessProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	if err := h.ProfileService.SaveProfile(r.Context(), profile); err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Profile saved successfully"})
}

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Authentication required")
		return
	}

	profile, err := h.ProfileService.GetProfile(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	// Ensure we don't expose internal user ID
	response := struct {
		*models.FitnessProfile
		UserID int `json:"-"`
	}{
		FitnessProfile: profile,
	}

	respondWithJSON(w, http.StatusOK, response)
}
