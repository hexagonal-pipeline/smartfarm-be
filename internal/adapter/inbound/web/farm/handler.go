package farm

import (
	"smartfarm-be/internal/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type Handler struct {
	useCase ports.FarmUseCase
}

func NewHandler(injector do.Injector) (*Handler, error) {
	useCase, err := do.Invoke[ports.FarmUseCase](injector)
	if err != nil {
		return nil, err
	}

	return &Handler{
		useCase: useCase,
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
	plots, err := h.useCase.ListAvailablePlots(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get available plots")
	}

	return c.JSON(NewFarmPlotListResponse(plots))
}
