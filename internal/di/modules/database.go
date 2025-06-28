package modules

import (
	"smartfarm-be/internal/adapter/outbound/db"
	"smartfarm-be/pkg/config"
	"smartfarm-be/pkg/database"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

func ProvideDatabase(injector do.Injector) {
	do.Provide(injector, func(i do.Injector) (*pgxpool.Pool, error) {
		postgresConfig, err := do.Invoke[*config.PostgresConfig](i)
		if err != nil {
			return nil, err
		}

		pool, err := database.NewPostgresPool(postgresConfig)
		if err != nil {
			return nil, err
		}

		return pool, nil
	})

	do.Provide(injector, func(i do.Injector) (db.Querier, error) {
		pool, err := do.Invoke[*pgxpool.Pool](i)
		if err != nil {
			return nil, err
		}

		return db.New(pool), nil
	})
}
