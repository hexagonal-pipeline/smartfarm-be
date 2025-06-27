package di

import (
	"smartfarm-be/internal/adapter/outbound/postgres"
	"smartfarm-be/internal/ports"

	"github.com/samber/do/v2"
)

func provideFarmRepository(injector do.Injector) {
	do.Provide(injector, func(i do.Injector) (ports.FarmRepository, error) {
		return postgres.NewFarmRepository(), nil
	})
}
