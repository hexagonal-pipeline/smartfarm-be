-- name: CreatePlantCard :one
INSERT INTO plant_cards (
    farm_plot_id,
    persona,
    image_url,
    video_url
) VALUES (
    $1, $2, $3, $4
)
RETURNING *; 