package di

import (
	"smartfarm-be/internal/ports"
	"smartfarm-be/internal/usecase"

	"github.com/samber/do/v2"
)

func provideFarmUseCase(injector do.Injector) {
	do.Provide(injector, func(i do.Injector) (ports.FarmUseCase, error) {
		repo, err := do.Invoke[ports.FarmRepository](i)
		if err != nil {
			return nil, err
		}
		return usecase.NewFarmService(repo), nil
	})
}
