package usecase

import (
	"context"
	"fmt"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports/outbound"

	"github.com/rs/zerolog/log"
	"github.com/samber/do/v2"
)

// PlantCardUsecase handles the logic for creating plant cards with persona, image, and video.
type PlantCardUsecase struct {
	farmRepo      outbound.FarmRepository
	plantCardRepo outbound.PlantCardRepository
	aiGen         outbound.AIGenerator
}

// NewPlantCardUsecase creates a new PlantCardUsecase.
func NewPlantCardUsecase(i do.Injector) (*PlantCardUsecase, error) {
	farmRepo, err := do.Invoke[outbound.FarmRepository](i)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke farm repository: %w", err)
	}

	plantCardRepo, err := do.Invoke[outbound.PlantCardRepository](i)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke plant card repository: %w", err)
	}

	aiGen, err := do.Invoke[outbound.AIGenerator](i)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke ai generator: %w", err)
	}

	return &PlantCardUsecase{
		farmRepo:      farmRepo,
		plantCardRepo: plantCardRepo,
		aiGen:         aiGen,
	}, nil
}

// GeneratePlantCard creates a new plant card for a given farm plot with persona, image, and video.
func (uc *PlantCardUsecase) GeneratePlantCard(ctx context.Context, farmPlotID int64) (*domain.PlantCard, error) {
	// 1. 농장 정보 조회
	farmPlot, err := uc.farmRepo.GetByID(ctx, farmPlotID)
	if err != nil {
		log.Error().Err(err).Int64("farm_plot_id", farmPlotID).Msg("failed to get farm plot")
		return nil, fmt.Errorf("failed to get farm plot: %w", err)
	}

	if farmPlot.PersonaPrompt == nil || *farmPlot.PersonaPrompt == "" {
		log.Error().Int64("farm_plot_id", farmPlotID).Msg("farm plot has no persona prompt")
		return nil, fmt.Errorf("farm plot %d has no persona prompt", farmPlotID)
	}

	personaPrompt := *farmPlot.PersonaPrompt

	// 2. 페르소나 생성
	persona, err := uc.aiGen.GeneratePersona(ctx, personaPrompt)
	if err != nil {
		log.Error().Err(err).Str("prompt", personaPrompt).Msg("failed to generate persona")
		return nil, fmt.Errorf("failed to generate persona: %w", err)
	}

	// 3. 이벤트 메시지 생성 (SNS 공유용)
	eventMessage, err := uc.aiGen.GenerateEventMessage(ctx, persona, "plant_card_creation")
	if err != nil {
		log.Error().Err(err).Str("persona", persona).Msg("failed to generate event message")
		return nil, fmt.Errorf("failed to generate event message: %w", err)
	}

	// 4. 이미지 생성
	imagePrompt := fmt.Sprintf("A cute and friendly character representing: %s. Style: cartoon, colorful, suitable for social media sharing", persona)
	imageURL, err := uc.aiGen.GenerateImage(ctx, imagePrompt)
	if err != nil {
		log.Error().Err(err).Str("prompt", imagePrompt).Msg("failed to generate image")
		return nil, fmt.Errorf("failed to generate image: %w", err)
	}

	// 5. Veo3를 이용한 쇼츠 비디오 생성
	videoPrompt := fmt.Sprintf("Create a short engaging video featuring this character: %s. The character says: %s", persona, eventMessage)
	videoURL, err := uc.aiGen.GenerateVideoFromPrompt(ctx, videoPrompt, imageURL)
	if err != nil {
		log.Error().Err(err).Str("prompt", videoPrompt).Msg("failed to generate video with Veo3")
		return nil, fmt.Errorf("failed to generate video: %w", err)
	}

	// 6. 플랜트카드 저장
	plantCard := domain.PlantCard{
		FarmPlotID:   farmPlotID,
		Persona:      persona,
		ImageURL:     imageURL,
		VideoURL:     videoURL,
		EventMessage: eventMessage,
	}

	result, err := uc.plantCardRepo.Create(ctx, &plantCard)
	if err != nil {
		log.Error().Err(err).Interface("plant_card", plantCard).Msg("failed to create plant card")
		return nil, fmt.Errorf("failed to create plant card: %w", err)
	}

	log.Info().Int64("plant_card_id", result.ID).Int64("farm_plot_id", farmPlotID).Msg("successfully created plant card")
	return result, nil
}

// GetPlantCardByID retrieves a plant card by its ID.
func (uc *PlantCardUsecase) GetPlantCardByID(ctx context.Context, id int64) (*domain.PlantCard, error) {
	plantCard, err := uc.plantCardRepo.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Int64("plant_card_id", id).Msg("failed to get plant card by ID")
		return nil, fmt.Errorf("failed to get plant card: %w", err)
	}

	log.Info().Int64("plant_card_id", id).Msg("successfully retrieved plant card")
	return plantCard, nil
}

// GetPlantCardsByFarmPlotID retrieves all plant cards for a specific farm plot.
func (uc *PlantCardUsecase) GetPlantCardsByFarmPlotID(ctx context.Context, farmPlotID int64) ([]domain.PlantCard, error) {
	plantCards, err := uc.plantCardRepo.GetByFarmPlotID(ctx, farmPlotID)
	if err != nil {
		log.Error().Err(err).Int64("farm_plot_id", farmPlotID).Msg("failed to get plant cards by farm plot ID")
		return nil, fmt.Errorf("failed to get plant cards: %w", err)
	}

	log.Info().Int64("farm_plot_id", farmPlotID).Int("count", len(plantCards)).Msg("successfully retrieved plant cards")
	return plantCards, nil
}
