package farm

import "github.com/gofiber/fiber/v2"

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/farms/plots/available", h.ListAvailablePlots)
}
