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
	db.Querier
}

func NewCommissionRepository(injector do.Injector) (outbound.CommissionRepository, error) {
	queries, err := do.Invoke[db.Querier](injector)
	if err != nil {
		return nil, err
	}
	return &CommissionRepository{
		Querier: queries,
	}, nil
}

func toDomainCommissionWork(w db.CommissionWork) *domain.CommissionWork {
	return &domain.CommissionWork{
		ID:                int64(w.ID),
		RequesterNickname: w.RequesterNickname,
		PlotID:            w.PlotID,
		TaskType:          w.TaskType,
		TaskDescription:   w.TaskDescription.String,
		Status:            w.Status,
		CreditCost:        w.CreditCost,
		RequestedAt:       w.RequestedAt.Time,
		CompletedAt:       &w.CompletedAt.Time,
	}
}

func toDomainCommissionWorkSlice(works []db.CommissionWork) []domain.CommissionWork {
	domainWorks := make([]domain.CommissionWork, len(works))
	for i, w := range works {
		domainWorks[i] = *toDomainCommissionWork(w)
	}
	return domainWorks
}

func (r *CommissionRepository) Create(ctx context.Context, arg domain.CommissionWork) (*domain.CommissionWork, error) {
	created, err := r.Querier.CreateCommissionWork(ctx, db.CreateCommissionWorkParams{
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

	return toDomainCommissionWork(created), nil
}

func (r *CommissionRepository) FindByID(ctx context.Context, id int64) (*domain.CommissionWork, error) {
	found, err := r.Querier.GetCommissionWork(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return toDomainCommissionWork(found), nil
}

func (r *CommissionRepository) ListByStatus(ctx context.Context, status string) ([]domain.CommissionWork, error) {
	works, err := r.Querier.ListCommissionWorksByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	return toDomainCommissionWorkSlice(works), nil
}

func (r *CommissionRepository) ListByRequester(ctx context.Context, requesterNickname string) ([]domain.CommissionWork, error) {
	works, err := r.Querier.ListCommissionWorksByRequester(ctx, requesterNickname)
	if err != nil {
		return nil, err
	}

	return toDomainCommissionWorkSlice(works), nil
}

func (r *CommissionRepository) ListByRequesterAndStatus(ctx context.Context, requesterNickname string, status string) ([]domain.CommissionWork, error) {
	works, err := r.Querier.ListCommissionWorksByRequesterAndStatus(ctx, db.ListCommissionWorksByRequesterAndStatusParams{
		RequesterNickname: requesterNickname,
		Status:            status,
	})
	if err != nil {
		return nil, err
	}

	return toDomainCommissionWorkSlice(works), nil
}
