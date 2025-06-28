package postgres

import (
	"context"

	"smartfarm-be/internal/adapter/outbound/db"
	"smartfarm-be/internal/domain"
	farmoutbound "smartfarm-be/internal/ports/outbound"

	"github.com/samber/do/v2"
)

type FarmRepository struct {
	querier db.Querier
}

func NewFarmRepository(injector do.Injector) (farmoutbound.FarmRepository, error) {
	querier, err := do.Invoke[db.Querier](injector)
	if err != nil {
		return nil, err
	}

	return &FarmRepository{querier: querier}, nil
}

func toDomainFarmPlot(p db.FarmPlot) domain.FarmPlot {
	return domain.FarmPlot{
		ID:          int64(p.ID),
		Name:        p.Name,
		IsAvailable: p.Status.String == "available" && p.Status.Valid,
	}
}

func (r *FarmRepository) ListAvailable(ctx context.Context) ([]domain.FarmPlot, error) {
	dbPlots, err := r.querier.ListAvailablePlots(ctx)
	if err != nil {
		return nil, err
	}

	domainPlots := make([]domain.FarmPlot, 0, len(dbPlots))
	for _, p := range dbPlots {
		domainPlots = append(domainPlots, toDomainFarmPlot(p))
	}

	return domainPlots, nil
}
