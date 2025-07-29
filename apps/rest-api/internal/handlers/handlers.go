package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"fmt"

	"rest-api/internal/middleware"
	"rest-api/internal/models"
	"rest-api/internal/services"
)

type Handlers struct {
	AuthService    *services.AuthService
	ProfileService *services.ProfileService
	AIService      *services.AIService
	HealthService  *services.HealthService
	MediaService   *services.MediaService
}

func NewHandlers(
	auth *services.AuthService,
	profile *services.ProfileService,
	ai *services.AIService,
	health *services.HealthService,
	media *services.MediaService,
) *Handlers {
	return &Handlers{
		AuthService:    auth,
		ProfileService: profile,
		AIService:      ai,
		HealthService:  health,
		MediaService:   media,
	}
}

func (h *Handlers) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.AuthMiddleware(h.AuthService.JWTSecret)(next).ServeHTTP(w, r)
	})
}

// Helper functions
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
	})
}

func handleServiceError(w http.ResponseWriter, err error) {
	if svcErr, ok := err.(services.ServiceError); ok {
		respondWithError(w, svcErr.Code, svcErr.Message)
	} else {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
	}
}

func ApkDownloadHandler(w http.ResponseWriter, r *http.Request) {
	apkPath := filepath.Join("downloads", "app-release.apk")

	// Open the file
	file, err := os.Open(apkPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get file info to check if it exists
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}

	// Set the headers
	w.Header().Set("Content-Disposition", "attachment; filename=app-release.apk")
	w.Header().Set("Content-Type", "application/vnd.android.package-archive")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Serve the file
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
