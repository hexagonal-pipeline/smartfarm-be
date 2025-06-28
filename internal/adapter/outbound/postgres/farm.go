package postgres

import (
	"context"
	"fmt"
	"smartfarm-be/internal/adapter/outbound/db"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports/outbound"

	"github.com/samber/do/v2"
)

type FarmRepository struct {
	q db.Querier
}

func NewFarmRepository(injector do.Injector) (outbound.FarmRepository, error) {
	querier, err := do.Invoke[db.Querier](injector)
	if err != nil {
		return nil, err
	}
	return &FarmRepository{q: querier}, nil
}

func toDomainFarmPlot(p db.FarmPlot) domain.FarmPlot {
	return domain.FarmPlot{
		ID:          int64(p.ID),
		Name:        p.Name,
		Location:    p.Location.String,
		SizeSqm:     p.SizeSqm,
		MonthlyRent: p.MonthlyRent,
		CropType:    p.CropType.String,
		Status:      p.Status.String,
		CreatedAt:   p.CreatedAt.Time,
	}
}

func toDomainFarmPlotFromRenterRow(p db.ListPlotsByRenterRow) domain.FarmPlot {
	return domain.FarmPlot{
		ID:       int64(p.ID),
		Name:     p.Name,
		Location: p.Location.String,
		SizeSqm:  p.SizeSqm,
		CropType: p.CropType.String,
		Status:   p.Status.String,
	}
}

func (r *FarmRepository) ListAvailable(ctx context.Context) ([]domain.FarmPlot, error) {
	dbPlots, err := r.q.ListAvailablePlots(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list available plots: %w", err)
	}

	plots := make([]domain.FarmPlot, len(dbPlots))
	for i, p := range dbPlots {
		plots[i] = toDomainFarmPlot(p)
	}

	return plots, nil
}

func (r *FarmRepository) ListByRenter(ctx context.Context, renterNickname string) ([]domain.FarmPlot, error) {
	dbPlots, err := r.q.ListPlotsByRenter(ctx, renterNickname)
	if err != nil {
		return nil, fmt.Errorf("failed to list plots by renter: %w", err)
	}

	plots := make([]domain.FarmPlot, len(dbPlots))
	for i, p := range dbPlots {
		plots[i] = toDomainFarmPlotFromRenterRow(p)
	}

	return plots, nil
}
