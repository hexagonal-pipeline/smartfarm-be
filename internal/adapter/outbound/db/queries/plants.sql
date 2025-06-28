-- name: CreatePlant :one
INSERT INTO plants (
    farm_plot_id,
    name,
    persona_context,
    image_url,
    description
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPlant :one
SELECT * FROM plants
WHERE id = $1 LIMIT 1;

-- name: ListPlantsByFarmPlotID :many
SELECT * FROM plants
WHERE farm_plot_id = $1
ORDER BY created_at DESC;

-- name: UpdatePlantInfo :one
UPDATE plants
SET
    persona_context = $2,
    image_url = $3,
    description = $4
WHERE id = $1
RETURNING *; 