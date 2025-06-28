package outbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

type CommissionRepository interface {
	Create(ctx context.Context, arg domain.CommissionWork) (*domain.CommissionWork, error)
	FindByID(ctx context.Context, id int64) (*domain.CommissionWork, error)
	ListByStatus(ctx context.Context, status string) ([]domain.CommissionWork, error)
	ListByRequester(ctx context.Context, requesterNickname string) ([]domain.CommissionWork, error)
	ListByRequesterAndStatus(ctx context.Context, requesterNickname string, status string) ([]domain.CommissionWork, error)
}
