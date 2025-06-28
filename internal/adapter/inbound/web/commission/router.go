package commission

import "github.com/gofiber/fiber/v2"

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/commissions", h.CreateCommissionWork)
}
