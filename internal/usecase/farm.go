package usecase

import (
	"context"
	"fmt"
	"smartfarm-be/internal/domain"
	farmoutbound "smartfarm-be/internal/ports/outbound"

	"github.com/samber/do/v2"
)

type FarmService struct {
	repo farmoutbound.FarmRepository
}

func NewFarmService(injector do.Injector) (*FarmService, error) {
	repo, err := do.Invoke[farmoutbound.FarmRepository](injector)
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

func (s *FarmService) ListMyPlots(ctx context.Context, nickname string) ([]domain.FarmPlot, error) {
	plots, err := s.repo.ListByRenter(ctx, nickname)
	if err != nil {
		return nil, fmt.Errorf("failed to list my plots: %w", err)
	}

	return plots, nil
}
