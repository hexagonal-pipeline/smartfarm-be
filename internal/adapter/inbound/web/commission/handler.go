package commission

import (
	"smartfarm-be/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type Handler struct {
	useCase *usecase.CommissionService
}

func NewHandler(injector do.Injector) (*Handler, error) {
	useCase, err := do.Invoke[*usecase.CommissionService](injector)
	if err != nil {
		return nil, err
	}

	return &Handler{
		useCase: useCase,
	}, nil
}

// CreateCommissionWork godoc
// @Summary      Create a new commission work
// @Description  register a new commission work for a plot
// @Tags         commissions
// @Accept       json
// @Produce      json
// @Param        commission_work  body   CreateCommissionWorkRequest  true  "Commission Work registration request"
// @Success      201  {object}  CommissionWorkResponse
// @Router       /commissions [post]
func (h *Handler) CreateCommissionWork(c *fiber.Ctx) error {
	var req CreateCommissionWorkRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	params := usecase.CreateCommissionWorkParams{
		RequesterNickname: req.Nickname,
		PlotID:            req.PlotID,
		TaskType:          req.TaskType,
		TaskDescription:   req.TaskDescription,
		CreditCost:        req.CreditCost,
	}

	createdWork, err := h.useCase.CreateCommissionWork(c.Context(), params)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create commission work")
	}

	return c.Status(fiber.StatusCreated).JSON(NewCommissionWorkResponse(createdWork))
}

// ListCommissionWorksByStatus godoc
// @Summary      List commission works by status
// @Description  get a list of commission works filtered by status
// @Tags         commissions
// @Accept       json
// @Produce      json
// @Param        status   query      string  true  "Status to filter by (e.g., requested, in_progress, completed, cancelled)"
// @Success      200  {object}  CommissionWorkListResponse
// @Router       /commissions [get]
func (h *Handler) ListCommissionWorksByStatus(c *fiber.Ctx) error {
	status := c.Query("status")
	if status == "" {
		return fiber.NewError(fiber.StatusBadRequest, "status query parameter is required")
	}

	works, err := h.useCase.ListCommissionWorksByStatus(c.Context(), status)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get commission works")
	}

	return c.JSON(NewCommissionWorkListResponse(works))
}
