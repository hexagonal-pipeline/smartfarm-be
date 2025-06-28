package modules

import (
	"smartfarm-be/pkg/config"

	"github.com/samber/do/v2"
)

func ProvideConfig(injector do.Injector) {
	do.Provide(injector, config.NewGeminiConfig)
	do.Provide(injector, config.NewPostgresConfig)
}
