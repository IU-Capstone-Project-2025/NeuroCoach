package handlers

import (
	"net/http"
	"rest-api/internal/models"
)

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Check database connection
	if err := h.HealthService.CheckDB(r.Context()); err != nil {
		respondWithError(w, http.StatusServiceUnavailable, "Database unavailable")
		return
	}

	// TODO: Add other health checks (e.g., AI service status)

	respondWithJSON(w, http.StatusOK, models.HealthCheckResponse{
		Status:  "healthy",
		Version: "1.0",
	})
}
