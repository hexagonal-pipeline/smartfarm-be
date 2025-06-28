package modules

import (
	"smartfarm-be/internal/adapter/outbound/postgres"

	"github.com/samber/do/v2"
)

func ProvideRepositories(injector do.Injector) {
	do.Provide(injector, postgres.NewFarmRepository)
	do.Provide(injector, postgres.NewCommissionRepository)
}
