package farmoutbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

// FarmRepository is the outbound port for farm-related database operations.
type FarmRepository interface {
	ListAvailable(ctx context.Context) ([]domain.FarmPlot, error)
}
