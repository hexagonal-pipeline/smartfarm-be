package di

import (
	"smartfarm-be/internal/di/modules"

	"github.com/samber/do/v2"
)

func InitializeInjector() (do.Injector, error) {
	injector := do.New()

	// Provide instances
	modules.ProvideConfig(injector)

	modules.ProvideDatabase(injector)
	modules.ProvideServer(injector)
	modules.ProvideRepositories(injector)
	modules.ProvideUseCases(injector)
	modules.ProvideHandlers(injector)

	return injector, nil
}
