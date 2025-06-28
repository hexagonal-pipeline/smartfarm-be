-- name: CreatePlantCard :one
INSERT INTO plant_cards (
    farm_plot_id,
    persona,
    image_url,
    video_url,
    event_message
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPlantCardByID :one
SELECT * FROM plant_cards WHERE id = $1;

-- name: GetPlantCardsByFarmPlotID :many
SELECT * FROM plant_cards WHERE farm_plot_id = $1 ORDER BY created_at DESC; 