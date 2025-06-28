package modules

import (
	"smartfarm-be/internal/adapter/outbound/googleai"

	"github.com/samber/do/v2"
)

func ProvideAdapters(injector do.Injector) {
	do.Provide(injector, googleai.NewGoogleAIGenerator)
}
