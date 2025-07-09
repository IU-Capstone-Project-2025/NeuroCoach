package services

import (
	"context"
	"net/http"

	"rest-api/internal/models"
	"rest-api/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaService struct {
	BaseService
}

func NewMediaService(repo repository.Repository, mongoRepo repository.MongoDBRep) *MediaService {
	return &MediaService{
		BaseService: BaseService{Repo: repo, MongoDBRepo: mongoRepo},
	}
}

func (s *MediaService) SaveExerciseMedia(ctx context.Context, req *models.ExerciseMediaRequest) error {
	// Convert exercise ID from string to ObjectID
	exerciseID, err := primitive.ObjectIDFromHex(req.ExerciseID)
	if err != nil {
		return NewServiceError(
			http.StatusBadRequest,
			"Invalid exercise ID format",
			err,
		)
	}

	// Create media object
	media := &models.ExerciseMedia{
		ExerciseID:  exerciseID,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Order:       req.Order,
	}

	// Save to database
	if err := s.MongoDBRepo.SaveExerciseMedia(ctx, media); err != nil {
		return NewServiceError(
			http.StatusInternalServerError,
			"Failed to save exercise media",
			err,
		)
	}

	return nil
}

func (s *MediaService) GetExerciseMedia(ctx context.Context, exerciseID string) ([]models.ExerciseMedia, error) {
	// Validate exercise ID format
	if _, err := primitive.ObjectIDFromHex(exerciseID); err != nil {
		return nil, NewServiceError(
			http.StatusBadRequest,
			"Invalid exercise ID format",
			err,
		)
	}

	// Get media from database
	media, err := s.MongoDBRepo.GetExerciseMedia(ctx, exerciseID)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to get exercise media",
			err,
		)
	}

	return media, nil
}

func (s *MediaService) DeleteExerciseMedia(ctx context.Context, mediaID string) error {
	// Validate media ID format
	if _, err := primitive.ObjectIDFromHex(mediaID); err != nil {
		return NewServiceError(
			http.StatusBadRequest,
			"Invalid media ID format",
			err,
		)
	}

	// Delete from database
	if err := s.MongoDBRepo.DeleteExerciseMedia(ctx, mediaID); err != nil {
		return NewServiceError(
			http.StatusInternalServerError,
			"Failed to delete exercise media",
			err,
		)
	}

	return nil
}