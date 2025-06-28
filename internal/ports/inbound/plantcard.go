package inbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

// PlantCardUsecase defines the interface for managing plant cards.
type PlantCardUsecase interface {
	GeneratePlantCard(ctx context.Context, farmPlotID int64) (*domain.PlantCard, error)
}
