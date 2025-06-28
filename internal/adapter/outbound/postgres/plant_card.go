package postgres

import (
	"context"
	"smartfarm-be/internal/adapter/outbound/db"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports/outbound"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/do/v2"
)

type PlantCardRepository struct {
	q db.Querier
}

func NewPlantCardRepository(injector do.Injector) (outbound.PlantCardRepository, error) {
	querier, err := do.Invoke[db.Querier](injector)
	if err != nil {
		return nil, err
	}
	return &PlantCardRepository{q: querier}, nil
}

func (r *PlantCardRepository) Create(ctx context.Context, card *domain.PlantCard) (*domain.PlantCard, error) {
	dbCard, err := r.q.CreatePlantCard(ctx, db.CreatePlantCardParams{
		FarmPlotID: int32(card.FarmPlotID),
		Persona:    card.Persona,
		ImageUrl:   pgtype.Text{String: card.ImageURL, Valid: card.ImageURL != ""},
		VideoUrl:   pgtype.Text{String: card.VideoURL, Valid: card.VideoURL != ""},
	})
	if err != nil {
		return nil, err
	}
	return &domain.PlantCard{
		ID:         int64(dbCard.ID),
		FarmPlotID: int64(dbCard.FarmPlotID),
		Persona:    dbCard.Persona,
		ImageURL:   dbCard.ImageUrl.String,
		VideoURL:   dbCard.VideoUrl.String,
		CreatedAt:  dbCard.CreatedAt.Time,
	}, nil
}
