package postgres

import (
	"context"
	"log"

	"smartfarm-be/internal/domain"
	"smartfarm-be/internal/ports"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

type FarmRepository struct {
	db *pgxpool.Pool
}

func NewFarmRepository(injector do.Injector) (ports.FarmRepository, error) {
	db, err := do.Invoke[*pgxpool.Pool](injector)
	if err != nil {
		return nil, err
	}

	return &FarmRepository{db: db}, nil
}

func (r *FarmRepository) ListAvailable(ctx context.Context) ([]domain.FarmPlot, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, status FROM farm_plots WHERE status = 'available'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plots []domain.FarmPlot
	for rows.Next() {
		var plot domain.FarmPlot
		var status string
		if err := rows.Scan(&plot.ID, &plot.Name, &status); err != nil {
			log.Printf("Error scanning farm plot: %v", err)
			continue
		}
		plot.IsAvailable = (status == "available")
		plots = append(plots, plot)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return plots, nil
}
