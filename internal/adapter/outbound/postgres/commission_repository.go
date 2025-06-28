package postgres

import (
	"context"
	"smartfarm-be/internal/adapter/outbound/db"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports/outbound"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/do/v2"
)

type CommissionRepository struct {
	*db.Queries
}

func NewCommissionRepository(injector do.Injector) (outbound.CommissionRepository, error) {
	queries, err := do.Invoke[*db.Queries](injector)
	if err != nil {
		return nil, err
	}
	return &CommissionRepository{
		Queries: queries,
	}, nil
}

func (r *CommissionRepository) Create(ctx context.Context, arg domain.CommissionWork) (*domain.CommissionWork, error) {
	created, err := r.Queries.CreateCommissionWork(ctx, db.CreateCommissionWorkParams{
		RequesterNickname: arg.RequesterNickname,
		PlotID:            arg.PlotID,
		TaskType:          arg.TaskType,
		TaskDescription: pgtype.Text{
			String: arg.TaskDescription,
			Valid:  true,
		},
		CreditCost: arg.CreditCost,
	})
	if err != nil {
		return nil, err
	}

	// TODO: Mapper 함수로 분리하는 것을 고려
	return &domain.CommissionWork{
		ID:                int64(created.ID),
		RequesterNickname: created.RequesterNickname,
		PlotID:            created.PlotID,
		TaskType:          created.TaskType,
		TaskDescription:   created.TaskDescription.String,
		Status:            created.Status,
		CreditCost:        created.CreditCost,
		RequestedAt:       created.RequestedAt.Time,
		// CompletedAt은 생성 시점이므로 nil
	}, nil
}

func (r *CommissionRepository) FindByID(ctx context.Context, id int64) (*domain.CommissionWork, error) {
	found, err := r.Queries.GetCommissionWork(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	// TODO: Mapper 함수로 분리하는 것을 고려
	return &domain.CommissionWork{
		ID:                int64(found.ID),
		RequesterNickname: found.RequesterNickname,
		PlotID:            found.PlotID,
		TaskType:          found.TaskType,
		TaskDescription:   found.TaskDescription.String,
		Status:            found.Status,
		CreditCost:        found.CreditCost,
		RequestedAt:       found.RequestedAt.Time,
		// CompletedAt 처리 필요
	}, nil
}
