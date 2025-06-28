package outbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

// FarmRepository is the outbound port for farm-related database operations.
type FarmRepository interface {
	ListAvailable(ctx context.Context) ([]domain.FarmPlot, error)
	ListByRenter(ctx context.Context, renterNickname string) ([]domain.FarmPlot, error)
	GetByID(ctx context.Context, id int64) (*domain.FarmPlot, error)
}
