package di

import (
	"smartfarm-be/internal/adapter/inbound/web/farm"
	"smartfarm-be/internal/ports"

	"github.com/samber/do/v2"
)

func provideFarmRouter(injector do.Injector) {
	do.Provide(injector, func(i do.Injector) (*farm.Router, error) {
		handler, err := do.Invoke[*farm.Handler](i)
		if err != nil {
			return nil, err
		}
		return farm.NewRouter(handler), nil
	})
}

func provideFarmHandler(injector do.Injector) {
	do.Provide(injector, func(i do.Injector) (*farm.Handler, error) {
		usecase, err := do.Invoke[ports.FarmUseCase](i)
		if err != nil {
			return nil, err
		}
		return farm.NewHandler(usecase), nil
	})
}
