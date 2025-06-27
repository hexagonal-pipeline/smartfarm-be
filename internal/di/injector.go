package di

import (
	"github.com/samber/do/v2"
)

func InitializeInjector() (do.Injector, error) {
	injector := do.New()

	// Provide instances
	provideApp(injector)
	provideFarmRepository(injector)
	provideFarmUseCase(injector)
	provideFarmHandler(injector)

	return injector, nil
}
