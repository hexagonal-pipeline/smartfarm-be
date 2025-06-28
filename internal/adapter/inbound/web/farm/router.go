package farm

import "github.com/gofiber/fiber/v2"

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/farms/plots/available", h.ListAvailablePlots)
	app.Get("/farms/my-plots", h.ListMyPlots)
	app.Post("/farms/plots/:id/plantcard", h.GeneratePlantCard)
	app.Get("/farms/plots/:id/plantcards", h.GetPlantCardsByFarmPlot)
}
