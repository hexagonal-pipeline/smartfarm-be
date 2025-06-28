package plantcard

import (
	"smartfarm-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type Handler struct {
	plantCardUsecase *usecase.PlantCardUsecase
}

func NewHandler(i do.Injector) (*Handler, error) {
	plantCardUsecase := do.MustInvoke[*usecase.PlantCardUsecase](i)
	return &Handler{
		plantCardUsecase: plantCardUsecase,
	}, nil
}

type CreatePlantCardRequest struct {
	FarmPlotID int64 `json:"farm_plot_id"`
}

type CreatePlantCardResponse struct {
	ID           int64  `json:"id"`
	FarmPlotID   int64  `json:"farm_plot_id"`
	Persona      string `json:"persona"`
	ImageURL     string `json:"image_url"`
	VideoURL     string `json:"video_url"`
	EventMessage string `json:"event_message"`
	CreatedAt    string `json:"created_at"`
}

// CreatePlantCard godoc
// @Summary      Create plant card
// @Description  Generate a plant card with persona, image, and video for a farm plot
// @Tags         plantcards
// @Accept       json
// @Produce      json
// @Param        request   body      CreatePlantCardRequest  true  "Plant card creation request"
// @Success      201  {object}   CreatePlantCardResponse
// @Router       /plantcards [post]
func (h *Handler) CreatePlantCard(c *fiber.Ctx) error {
	var req CreatePlantCardRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if req.FarmPlotID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "farm_plot_id is required")
	}

	plantCard, err := h.plantCardUsecase.GeneratePlantCard(c.Context(), req.FarmPlotID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create plant card")
	}

	response := CreatePlantCardResponse{
		ID:           plantCard.ID,
		FarmPlotID:   plantCard.FarmPlotID,
		Persona:      plantCard.Persona,
		ImageURL:     plantCard.ImageURL,
		VideoURL:     plantCard.VideoURL,
		EventMessage: plantCard.EventMessage,
		CreatedAt:    plantCard.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetPlantCard godoc
// @Summary      Get plant card by ID
// @Description  Get plant card data by ID
// @Tags         plantcards
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Plant Card ID"
// @Success      200  {object}   CreatePlantCardResponse
// @Router       /plantcards/{id} [get]
func (h *Handler) GetPlantCard(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid plant card ID")
	}

	plantCard, err := h.plantCardUsecase.GetPlantCardByID(c.Context(), int64(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Plant card not found")
	}

	response := CreatePlantCardResponse{
		ID:           plantCard.ID,
		FarmPlotID:   plantCard.FarmPlotID,
		Persona:      plantCard.Persona,
		ImageURL:     plantCard.ImageURL,
		VideoURL:     plantCard.VideoURL,
		EventMessage: plantCard.EventMessage,
		CreatedAt:    plantCard.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return c.JSON(response)
}

// SharePlantCard godoc
// @Summary      Get plant card for SNS sharing
// @Description  Get plant card data formatted for social media sharing
// @Tags         plantcards
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Plant Card ID"
// @Success      200  {object}   CreatePlantCardResponse
// @Router       /plantcards/{id}/share [get]
func (h *Handler) SharePlantCard(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid plant card ID")
	}

	plantCard, err := h.plantCardUsecase.GetPlantCardByID(c.Context(), int64(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Plant card not found")
	}

	// SNS 공유용 특별한 포맷 (현재는 동일하지만 추후 확장 가능)
	response := CreatePlantCardResponse{
		ID:           plantCard.ID,
		FarmPlotID:   plantCard.FarmPlotID,
		Persona:      plantCard.Persona,
		ImageURL:     plantCard.ImageURL,
		VideoURL:     plantCard.VideoURL,
		EventMessage: plantCard.EventMessage,
		CreatedAt:    plantCard.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return c.JSON(response)
}
