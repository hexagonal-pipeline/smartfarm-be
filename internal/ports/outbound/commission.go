package outbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

type CommissionRepository interface {
	Create(ctx context.Context, arg domain.CommissionWork) (*domain.CommissionWork, error)
	FindByID(ctx context.Context, id int64) (*domain.CommissionWork, error)
}
