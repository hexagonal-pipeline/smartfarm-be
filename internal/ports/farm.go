package ports

import (
	"context"
	"smartfarm-be/internal/domain"
)

// FarmUseCase is the inbound port for farm-related operations.
type FarmUseCase interface {
	ListAvailablePlots(ctx context.Context) ([]domain.FarmPlot, error)
}

// FarmRepository is the outbound port for farm-related database operations.
type FarmRepository interface {
	ListAvailable(ctx context.Context) ([]domain.FarmPlot, error)
}
