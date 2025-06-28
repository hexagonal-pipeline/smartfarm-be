package plantcard

import "github.com/gofiber/fiber/v2"

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/videogen", h.GeneratePlantCard)
	app.Post("/plantcards", h.CreatePlantCard)
	app.Get("/plantcards/:id", h.GetPlantCard)
	app.Get("/plantcards/:id/share", h.SharePlantCard)
}
