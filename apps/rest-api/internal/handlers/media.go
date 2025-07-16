package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"rest-api/internal/models"
)

func (h *Handlers) SaveExerciseMedia(w http.ResponseWriter, r *http.Request) {
	var req models.ExerciseMediaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	if err := h.MediaService.SaveExerciseMedia(r.Context(), &req); err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message": "Exercise media saved successfully",
	})
}

func (h *Handlers) GetAllExerciseMedia(w http.ResponseWriter, r *http.Request) {
	media, err := h.MediaService.GetAllExerciseMedia(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, media)
}

func (h *Handlers) DeleteExerciseMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID := vars["media_id"]

	if err := h.MediaService.DeleteExerciseMedia(r.Context(), mediaID); err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Exercise media deleted successfully",
	})
}