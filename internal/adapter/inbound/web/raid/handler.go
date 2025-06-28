package raid

import (
	"smartfarm-be/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type Handler struct {
	useCase *usecase.RaidService
}

func NewHandler(injector do.Injector) (*Handler, error) {
	useCase, err := do.Invoke[*usecase.RaidService](injector)
	if err != nil {
		return nil, err
	}

	return &Handler{
		useCase: useCase,
	}, nil
}

// ListOpenRaids godoc
// @Summary      List open raids
// @Description  Get a list of all open raids
// @Tags         raids
// @Accept       json
// @Produce      json
// @Success      200  {array}   RaidResponse
// @Router       /raids [get]
func (h *Handler) ListOpenRaids(c *fiber.Ctx) error {
	raids, err := h.useCase.ListOpenRaids(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get raids")
	}

	return c.JSON(NewRaidListResponse(raids))
}

// GetRaidDetails godoc
// @Summary      Get raid details
// @Description  Get raid details by ID
// @Tags         raids
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Raid ID"
// @Success      200  {object}   RaidResponse
// @Router       /raids/{id} [get]
func (h *Handler) GetRaidDetails(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid raid ID")
	}

	raid, err := h.useCase.GetRaidDetails(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Raid not found")
	}

	return c.JSON(NewRaidResponse(*raid))
}

// JoinRaid godoc
// @Summary      Join raid
// @Description  Join a raid with specified quantity
// @Tags         raids
// @Accept       json
// @Produce      json
// @Param        id      path      int                 true  "Raid ID"
// @Param        request body      JoinRaidRequest     true  "Join raid request"
// @Success      201  {object}   RaidParticipationResponse
// @Router       /raids/{id}/join [post]
func (h *Handler) JoinRaid(c *fiber.Ctx) error {
	idStr := c.Params("id")
	raidID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid raid ID")
	}

	var req JoinRaidRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Nickname == "" || req.Quantity <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Nickname and quantity are required")
	}

	participation, err := h.useCase.JoinRaid(c.Context(), raidID, req.Nickname, req.Quantity)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to join raid")
	}

	return c.Status(fiber.StatusCreated).JSON(NewRaidParticipationResponse(*participation))
}

// GetRaidParticipations godoc
// @Summary      Get raid participations
// @Description  Get all participations for a specific raid
// @Tags         raids
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Raid ID"
// @Success      200  {array}   RaidParticipationResponse
// @Router       /raids/{id}/participations [get]
func (h *Handler) GetRaidParticipations(c *fiber.Ctx) error {
	idStr := c.Params("id")
	raidID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid raid ID")
	}

	participations, err := h.useCase.GetRaidParticipations(c.Context(), raidID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get participations")
	}

	return c.JSON(NewRaidParticipationListResponse(participations))
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	raids := app.Group("/raids")
	raids.Get("/", h.ListOpenRaids)
	raids.Get("/:id", h.GetRaidDetails)
	raids.Post("/:id/join", h.JoinRaid)
	raids.Get("/:id/participations", h.GetRaidParticipations)
}
