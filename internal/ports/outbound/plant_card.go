package outbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

type PlantCardRepository interface {
	Create(ctx context.Context, params *domain.PlantCard) (*domain.PlantCard, error)
	GetByID(ctx context.Context, id int64) (*domain.PlantCard, error)
	GetByFarmPlotID(ctx context.Context, farmPlotID int64) ([]domain.PlantCard, error)
}
