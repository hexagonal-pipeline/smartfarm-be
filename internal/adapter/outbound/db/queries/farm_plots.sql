-- name: ListAvailablePlots :many
SELECT * FROM farm_plots
WHERE status = 'available'
ORDER BY monthly_rent ASC;

-- name: GetPlot :one
SELECT * FROM farm_plots
WHERE id = $1;

-- name: ListPlotsByCrop :many
SELECT * FROM farm_plots
WHERE crop_type = $1 AND status = 'available'
ORDER BY monthly_rent ASC;

-- name: UpdatePlotStatus :one
UPDATE farm_plots
SET status = $2
WHERE id = $1
RETURNING *; 