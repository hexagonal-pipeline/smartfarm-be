package usecase

import (
	"context"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports"
)

type FarmService struct {
	repo ports.FarmRepository
}

func NewFarmService(repo ports.FarmRepository) *FarmService {
	return &FarmService{
		repo: repo,
	}
}

func (s *FarmService) ListAvailablePlots(ctx context.Context) ([]domain.FarmPlot, error) {
	// Here you can add any business logic before or after fetching the data.
	// For now, it just calls the repository.
	return s.repo.ListAvailable(ctx)
}
