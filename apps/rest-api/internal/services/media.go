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
	// Create media object
	media := &models.ExerciseMedia{
		Name:        req.Name,
		ImageURL:    req.ImageURL,
		Description: req.Description,
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

func (s *MediaService) GetAllExerciseMedia(ctx context.Context) ([]models.ExerciseMedia, error) {
	// Get all media from database
	media, err := s.MongoDBRepo.GetAllExerciseMedia(ctx)
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
