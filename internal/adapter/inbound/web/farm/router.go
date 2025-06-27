package farm

import "github.com/gofiber/fiber/v2"

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{
		handler: handler,
	}
}

func (r *Router) RegisterRoutes(app fiber.Router) {
	app.Get("/farms/plots/available", r.handler.ListAvailablePlots)
}
