package modules

import (
	"smartfarm-be/internal/adapter/inbound/web/commission"
	"smartfarm-be/internal/adapter/inbound/web/farm"
	"smartfarm-be/internal/adapter/inbound/web/plantcard"

	"github.com/samber/do/v2"
)

func ProvideHandlers(injector do.Injector) {
	do.Provide(injector, farm.NewHandler)
	do.Provide(injector, commission.NewHandler)
	do.Provide(injector, plantcard.NewHandler)
}
