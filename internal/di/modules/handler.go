package modules

import (
	"smartfarm-be/internal/adapter/inbound/web/farm"

	"github.com/samber/do/v2"
)

func ProvideHandlers(injector do.Injector) {
	do.Provide(injector, farm.NewHandler)
}
