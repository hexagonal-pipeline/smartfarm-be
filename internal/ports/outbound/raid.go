package outbound

import (
	"context"
	"smartfarm-be/internal/domain"
)

type RaidRepository interface {
	ListOpenRaids(ctx context.Context) ([]domain.Raid, error)
	GetRaidDetails(ctx context.Context, id int64) (*domain.Raid, error)
	JoinRaid(ctx context.Context, raidID int64, nickname string, quantity int32, expectedRevenue int32) (*domain.RaidParticipation, error)
	GetRaidParticipations(ctx context.Context, raidID int64) ([]domain.RaidParticipation, error)
}
