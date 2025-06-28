package modules

import (
	"smartfarm-be/internal/adapter/outbound/postgres"

	"github.com/samber/do/v2"
)

func ProvideRepositories(injector do.Injector) {
	// Repositories
	do.Provide(injector, postgres.NewFarmRepository)
	do.Provide(injector, postgres.NewCommissionRepository)
	do.Provide(injector, postgres.NewPlantCardRepository)
	do.Provide(injector, postgres.NewRaidRepository)
}
