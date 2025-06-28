package outbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

type PlantCardRepository interface {
	Create(ctx context.Context, params *domain.PlantCard) (*domain.PlantCard, error)
}
