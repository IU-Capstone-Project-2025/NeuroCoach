package services

import (
	"context"
	"net/http"

	"rest-api/internal/models"
	"rest-api/internal/repository"
)

type ProfileService struct {
	BaseService
}

func NewProfileService(repo repository.Repository) *ProfileService {
	return &ProfileService{
		BaseService: BaseService{Repo: repo},
	}
}

func (s *ProfileService) SaveProfile(ctx context.Context, profile models.FitnessProfile) error {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	profile.UserID = userID
	if err := s.Repo.SaveFitnessProfile(ctx, userID, &profile); err != nil {
		return NewServiceError(
			http.StatusInternalServerError, 
			"Failed to save profile", 
			err,
		)
	}
	return nil
}

func (s *ProfileService) GetProfile(ctx context.Context) (*models.FitnessProfile, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	profile, err := s.Repo.GetFitnessProfile(ctx, userID)
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, NewServiceError(
				http.StatusNotFound, 
				"Profile not found", 
				err,
			)
		}
		return nil, NewServiceError(
			http.StatusInternalServerError, 
			"Failed to get profile", 
			err,
		)
	}
	return profile, nil
}