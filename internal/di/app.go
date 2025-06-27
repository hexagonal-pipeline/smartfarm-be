package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

func provideApp(injector do.Injector) {
	do.Provide(injector, func(i do.Injector) (*fiber.App, error) {
		return fiber.New(), nil
	})
}
