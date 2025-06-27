package usecase

import (
	"context"
	"fmt"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports"

	"github.com/samber/do/v2"
)

type FarmService struct {
	repo ports.FarmRepository
}

func NewFarmService(injector do.Injector) (ports.FarmUseCase, error) {
	repo, err := do.Invoke[ports.FarmRepository](injector)
	if err != nil {
		return nil, err
	}

	return &FarmService{repo: repo}, nil
}

func (s *FarmService) ListAvailablePlots(ctx context.Context) ([]domain.FarmPlot, error) {
	plots, err := s.repo.ListAvailable(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list available plots: %w", err)
	}

	return plots, nil
}
