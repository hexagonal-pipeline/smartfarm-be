package farm

import (
	"smartfarm-be/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase ports.FarmUseCase
}

func NewHandler(useCase ports.FarmUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
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
	plots, err := h.useCase.ListAvailablePlots(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get available plots")
	}

	return c.JSON(NewFarmPlotListResponse(plots))
}
