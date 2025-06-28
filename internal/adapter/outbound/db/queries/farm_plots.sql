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

-- name: UpdateFarmPlotStatus :one
UPDATE farm_plots
SET status = $2
WHERE id = $1
RETURNING *;

-- name: ListPlotsByRenter :many
SELECT
    p.id,
    p.name,
    p.location,
    p.size_sqm,
    p.crop_type,
    p.status,
    r.start_date,
    r.end_date
FROM
    farm_plots p
JOIN
    rentals r ON p.id = r.plot_id
WHERE
    r.renter_nickname = $1 AND r.status = 'active'
ORDER BY
    r.start_date DESC; 