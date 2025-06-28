package usecase

import (
	"context"
	"fmt"
	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports/inbound"
	"smartfarm-be/internal/ports/outbound"

	"github.com/samber/do/v2"
)

// PlantCardUsecase handles the logic for creating plant cards.
type PlantCardUsecase struct {
	farmRepo outbound.FarmRepository
	aiGen    outbound.AIGenerator
}

// NewPlantCardUsecase creates a new PlantCardUsecase.
func NewPlantCardUsecase(i do.Injector) (inbound.PlantCardUsecase, error) {
	farmRepo, err := do.Invoke[outbound.FarmRepository](i)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke farm repository: %w", err)
	}

	aiGen, err := do.Invoke[outbound.AIGenerator](i)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke ai generator: %w", err)
	}

	return &PlantCardUsecase{
		farmRepo: farmRepo,
		aiGen:    aiGen,
	}, nil
}

// GeneratePlantCard creates a new plant card for a given farm plot.
func (uc *PlantCardUsecase) GeneratePlantCard(ctx context.Context, farmPlotID int64) (*domain.PlantCard, error) {
	farmPlot, err := uc.farmRepo.GetByID(ctx, farmPlotID)
	if err != nil {
		return nil, fmt.Errorf("failed to get farm plot: %w", err)
	}
	if farmPlot.PersonaPrompt == nil || *farmPlot.PersonaPrompt == "" {
		return nil, fmt.Errorf("farm plot %d has no persona prompt", farmPlotID)
	}
	personaPrompt := *farmPlot.PersonaPrompt

	persona, err := uc.aiGen.GeneratePersona(ctx, personaPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate persona: %w", err)
	}

	introMessage, err := uc.aiGen.GenerateEventMessage(ctx, persona, "introduction")
	if err != nil {
		return nil, fmt.Errorf("failed to generate introduction message: %w", err)
	}

	imagePrompt := fmt.Sprintf("A cute drawing of a character: %s", introMessage)
	imageURL, err := uc.aiGen.GenerateImage(ctx, imagePrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate image: %w", err)
	}

	videoURL, err := uc.aiGen.GenerateVideo(ctx, persona, imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to generate video: %w", err)
	}

	plantCard := &domain.PlantCard{
		FarmPlotID: farmPlotID,
		Persona:    persona,
		ImageURL:   imageURL,
		VideoURL:   videoURL,
	}

	return plantCard, nil
}
