package usecase

import (
	"context"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports/outbound"

	"github.com/rs/zerolog/log"
	"github.com/samber/do/v2"
)

type CommissionService struct {
	repo outbound.CommissionRepository
}

func NewCommissionService(injector do.Injector) (*CommissionService, error) {
	repo, err := do.Invoke[outbound.CommissionRepository](injector)
	if err != nil {
		return nil, err
	}
	return &CommissionService{
		repo: repo,
	}, nil
}

type CreateCommissionWorkParams struct {
	RequesterNickname string
	PlotID            int32
	TaskType          string
	TaskDescription   string
	CreditCost        int32
}

func (s *CommissionService) CreateCommissionWork(ctx context.Context, params CreateCommissionWorkParams) (*domain.CommissionWork, error) {
	// TODO: 여기에 비즈니스 로직 추가 (예: 사용자의 크레딧 확인 및 차감)

	newWork := domain.CommissionWork{
		RequesterNickname: params.RequesterNickname,
		PlotID:            params.PlotID,
		TaskType:          params.TaskType,
		TaskDescription:   params.TaskDescription,
		CreditCost:        params.CreditCost,
	}

	createdWork, err := s.repo.Create(ctx, newWork)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create commission work")
		return nil, err
	}

	return createdWork, nil
}

func (s *CommissionService) ListCommissionWorksByStatus(ctx context.Context, status string) ([]domain.CommissionWork, error) {
	works, err := s.repo.ListByStatus(ctx, status)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list commission works by status")
		return nil, err
	}

	return works, nil
}
