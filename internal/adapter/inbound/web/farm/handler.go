package farm

import (
	"smartfarm-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type Handler struct {
	farmService      *usecase.FarmService
	plantCardUsecase *usecase.PlantCardUsecase
}

func NewHandler(i do.Injector) (*Handler, error) {
	farmService := do.MustInvoke[*usecase.FarmService](i)
	plantCardUsecase := do.MustInvoke[*usecase.PlantCardUsecase](i)
	return &Handler{
		farmService:      farmService,
		plantCardUsecase: plantCardUsecase,
	}, nil
}

// ListAvailablePlots godoc
// @Summary      List available farm plots
// @Description  get a list of all farm plots that are available
// @Tags         farms
// @Accept       json
// @Produce      json
// @Success      200  {array}   FarmPlotResponse
// @Router       /farms/plots/available [get]
func (h *Handler) ListAvailablePlots(c *fiber.Ctx) error {
	plots, err := h.farmService.ListAvailablePlots(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get available plots")
	}

	return c.JSON(NewFarmPlotListResponse(plots))
}

// ListMyPlots godoc
// @Summary      List my farm plots
// @Description  get a list of all farm plots rented by a user
// @Tags         farms
// @Accept       json
// @Produce      json
// @Param        nickname   query      string  true  "User nickname"
// @Success      200  {array}   FarmPlotResponse
// @Router       /farms/my-plots [get]
func (h *Handler) ListMyPlots(c *fiber.Ctx) error {
	nickname := c.Query("nickname")
	if nickname == "" {
		return fiber.NewError(fiber.StatusBadRequest, "nickname query parameter is required")
	}

	plots, err := h.farmService.ListMyPlots(c.Context(), nickname)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get my plots")
	}

	return c.JSON(NewFarmPlotListResponse(plots))
}

// GeneratePlantCard godoc
// @Summary      Generate plant card for farm plot
// @Description  Generate a plant card with persona, image, and video for a farm plot
// @Tags         farms
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Farm Plot ID"
// @Success      201  {object}   map[string]interface{}
// @Router       /farms/plots/{id}/plantcard [post]
func (h *Handler) GeneratePlantCard(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid plot ID")
	}

	plantCard, err := h.plantCardUsecase.GeneratePlantCard(ctx.Context(), int64(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate plant card")
	}

	return ctx.Status(fiber.StatusCreated).JSON(plantCard)
}

// GetPlantCardsByFarmPlot godoc
// @Summary      Get plant cards for farm plot
// @Description  Get all plant cards created for a specific farm plot
// @Tags         farms
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Farm Plot ID"
// @Success      200  {array}   map[string]interface{}
// @Router       /farms/plots/{id}/plantcards [get]
func (h *Handler) GetPlantCardsByFarmPlot(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid plot ID")
	}

	plantCards, err := h.plantCardUsecase.GetPlantCardsByFarmPlotID(ctx.Context(), int64(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get plant cards")
	}

	return ctx.JSON(plantCards)
}
