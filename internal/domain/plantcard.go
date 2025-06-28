package domain

import "time"

// PlantCard represents the generated social media content for a farm plot.
type PlantCard struct {
	ID           int64     `json:"id"`
	FarmPlotID   int64     `json:"farm_plot_id"`
	Persona      string    `json:"persona"`
	ImageURL     string    `json:"image_url"`
	VideoURL     string    `json:"video_url"`
	EventMessage string    `json:"event_message"`
	CreatedAt    time.Time `json:"created_at"`
}
