package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"rest-api/internal/models"
)

// SaveExerciseMedia godoc
// @Summary Save exercise media
// @Description Save exercise media with image URL and description
// @Tags media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.ExerciseMediaRequest true "Exercise media data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/media [post]
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

// GetAllExerciseMedia godoc
// @Summary Get all exercise media
// @Description Retrieve all exercise media entries
// @Tags media
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.ExerciseMedia
// @Failure 401 {object} models.ErrorResponse
// @Router /api/media [get]
func (h *Handlers) GetAllExerciseMedia(w http.ResponseWriter, r *http.Request) {
	media, err := h.MediaService.GetAllExerciseMedia(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, media)
}

// DeleteExerciseMedia godoc
// @Summary Delete exercise media
// @Description Delete exercise media by ID
// @Tags media
// @Produce json
// @Security BearerAuth
// @Param media_id path string true "Media ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/media/{media_id} [delete]
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
