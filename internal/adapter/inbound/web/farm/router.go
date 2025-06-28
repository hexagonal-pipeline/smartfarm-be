package farm

import "github.com/gofiber/fiber/v2"

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/farms/plots/available", h.ListAvailablePlots)
	app.Get("/farms/my-plots", h.ListMyPlots)
	app.Get("/farms/plots/:id/plantcards", h.GetPlantCardsByFarmPlot)
}
