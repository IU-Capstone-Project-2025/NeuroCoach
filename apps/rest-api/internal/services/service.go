package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"rest-api/internal/middleware"
	"rest-api/internal/repository"
)

type ServiceError struct {
	Code    int
	Message string
	Err     error
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func NewServiceError(code int, message string, err error) ServiceError {
	return ServiceError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

type BaseService struct {
	Repo        repository.Repository
	MongoDBRepo repository.MongoDBRep
}

func (s *BaseService) GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return 0, NewServiceError(
			http.StatusInternalServerError,
			"User ID missing in context",
			errors.New("user ID not found in request context"),
		)
	}
	return userID, nil
}
