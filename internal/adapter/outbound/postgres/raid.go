package postgres

import (
	"context"
	"fmt"
	"smartfarm-be/internal/adapter/outbound/db"
	"smartfarm-be/internal/domain"
	outbound "smartfarm-be/internal/ports/outbound"

	"github.com/samber/do/v2"
)

type raidRepository struct {
	queries db.Querier
}

func NewRaidRepository(injector do.Injector) (outbound.RaidRepository, error) {
	queries, err := do.Invoke[db.Querier](injector)
	if err != nil {
		return nil, fmt.Errorf("failed to get queries: %w", err)
	}

	return &raidRepository{
		queries: queries,
	}, nil
}

func (r *raidRepository) ListOpenRaids(ctx context.Context) ([]domain.Raid, error) {
	rows, err := r.queries.ListOpenRaids(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list open raids: %w", err)
	}

	raids := make([]domain.Raid, len(rows))
	for i, row := range rows {
		raids[i] = domain.Raid{
			ID:               int64(row.ID),
			Title:            row.Title,
			Description:      row.Description.String,
			CropType:         row.CropType,
			TargetQuantity:   row.TargetQuantity,
			MinParticipation: row.MinParticipation,
			MaxParticipation: row.MaxParticipation,
			PricePerKg:       row.PricePerKg,
			Deadline:         row.Deadline.Time,
			Status:           row.Status.String,
			CreatorNickname:  row.CreatorNickname,
			CreatedAt:        row.CreatedAt.Time,
			CurrentQuantity:  row.CurrentQuantity.(int64),
			ParticipantCount: row.ParticipantCount,
		}
	}

	return raids, nil
}

func (r *raidRepository) GetRaidDetails(ctx context.Context, id int64) (*domain.Raid, error) {
	row, err := r.queries.GetRaidDetails(ctx, int32(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get raid details: %w", err)
	}

	raid := &domain.Raid{
		ID:               int64(row.ID),
		Title:            row.Title,
		Description:      row.Description.String,
		CropType:         row.CropType,
		TargetQuantity:   row.TargetQuantity,
		MinParticipation: row.MinParticipation,
		MaxParticipation: row.MaxParticipation,
		PricePerKg:       row.PricePerKg,
		Deadline:         row.Deadline.Time,
		Status:           row.Status.String,
		CreatorNickname:  row.CreatorNickname,
		CreatedAt:        row.CreatedAt.Time,
		CurrentQuantity:  row.CurrentQuantity.(int64),
		ParticipantCount: row.ParticipantCount,
	}

	return raid, nil
}

func (r *raidRepository) JoinRaid(ctx context.Context, raidID int64, nickname string, quantity int32, expectedRevenue int32) (*domain.RaidParticipation, error) {
	participation, err := r.queries.JoinRaid(ctx, db.JoinRaidParams{
		RaidID:              int32(raidID),
		ParticipantNickname: nickname,
		Quantity:            quantity,
		ExpectedRevenue:     expectedRevenue,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to join raid: %w", err)
	}

	result := &domain.RaidParticipation{
		ID:                  int64(participation.ID),
		RaidID:              participation.RaidID,
		ParticipantNickname: participation.ParticipantNickname,
		Quantity:            participation.Quantity,
		ExpectedRevenue:     participation.ExpectedRevenue,
		Status:              participation.Status.String,
		CreatedAt:           participation.CreatedAt.Time,
	}

	return result, nil
}

func (r *raidRepository) GetRaidParticipations(ctx context.Context, raidID int64) ([]domain.RaidParticipation, error) {
	rows, err := r.queries.GetRaidParticipations(ctx, int32(raidID))
	if err != nil {
		return nil, fmt.Errorf("failed to get raid participations: %w", err)
	}

	participations := make([]domain.RaidParticipation, len(rows))
	for i, row := range rows {
		participations[i] = domain.RaidParticipation{
			ID:                  int64(row.ID),
			RaidID:              row.RaidID,
			ParticipantNickname: row.ParticipantNickname,
			Quantity:            row.Quantity,
			ExpectedRevenue:     row.ExpectedRevenue,
			Status:              row.Status.String,
			CreatedAt:           row.CreatedAt.Time,
		}
	}

	return participations, nil
}
