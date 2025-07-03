package services

import (
	"context"
	"rest-api/internal/repository"
)

type HealthService struct {
	BaseService
}

func NewHealthService(repo repository.Repository) *HealthService {
	return &HealthService{
		BaseService: BaseService{Repo: repo},
	}
}

func (s *HealthService) CheckDB(ctx context.Context) error {
	return s.Repo.Ping(ctx)
}
