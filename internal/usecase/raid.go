package usecase

import (
	"context"
	"fmt"
	"smartfarm-be/internal/domain"
	outbound "smartfarm-be/internal/ports/outbound"

	"github.com/samber/do/v2"
)

type RaidService struct {
	repo outbound.RaidRepository
}

func NewRaidService(injector do.Injector) (*RaidService, error) {
	repo, err := do.Invoke[outbound.RaidRepository](injector)
	if err != nil {
		return nil, err
	}

	return &RaidService{repo: repo}, nil
}

func (s *RaidService) ListOpenRaids(ctx context.Context) ([]domain.Raid, error) {
	raids, err := s.repo.ListOpenRaids(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list open raids: %w", err)
	}

	return raids, nil
}

func (s *RaidService) GetRaidDetails(ctx context.Context, id int64) (*domain.Raid, error) {
	raid, err := s.repo.GetRaidDetails(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get raid details: %w", err)
	}

	return raid, nil
}

func (s *RaidService) JoinRaid(ctx context.Context, raidID int64, nickname string, quantity int32) (*domain.RaidParticipation, error) {
	// 간단한 예상 수익 계산
	raid, err := s.repo.GetRaidDetails(ctx, raidID)
	if err != nil {
		return nil, fmt.Errorf("failed to get raid details: %w", err)
	}

	expectedRevenue := int32(quantity) * raid.PricePerKg

	participation, err := s.repo.JoinRaid(ctx, raidID, nickname, quantity, expectedRevenue)
	if err != nil {
		return nil, fmt.Errorf("failed to join raid: %w", err)
	}

	return participation, nil
}

func (s *RaidService) GetRaidParticipations(ctx context.Context, raidID int64) ([]domain.RaidParticipation, error) {
	participations, err := s.repo.GetRaidParticipations(ctx, raidID)
	if err != nil {
		return nil, fmt.Errorf("failed to get raid participations: %w", err)
	}

	return participations, nil
}
