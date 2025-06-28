package modules

import (
	"smartfarm-be/internal/usecase"

	"github.com/samber/do/v2"
)

func ProvideUseCases(injector do.Injector) {
	do.Provide(injector, usecase.NewFarmService)
	do.Provide(injector, usecase.NewCommissionService)
	do.Provide(injector, usecase.NewPlantCardUsecase)
	do.Provide(injector, usecase.NewRaidService)
}
