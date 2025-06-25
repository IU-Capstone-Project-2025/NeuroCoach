package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"rest-api/internal/models"
)

var validate = validator.New()

func ValidateRequest[T any](next func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request T

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := validate.Struct(request); err != nil {
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				errors[err.Field()] = err.Tag()
			}
			respondWithValidationError(w, errors)
			return
		}

		next(w, r, request)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
	})
}

func respondWithValidationError(w http.ResponseWriter, errors map[string]string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error:   "Validation failed",
		Message: "Invalid request parameters",
	})
}