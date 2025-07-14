package handlers

import (
	"net/http"
	"rest-api/internal/models"
)

// HealthCheck godoc
// @Summary Health check
// @Description Check API and database health
// @Tags health
// @Produce json
// @Success 200 {object} models.HealthCheckResponse
// @Failure 503 {object} models.ErrorResponse
// @Router /health [get]
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
