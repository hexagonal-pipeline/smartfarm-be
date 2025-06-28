package domain

import "time"

// PlantCard represents the generated social media content for a farm plot.
type PlantCard struct {
	ID         int64
	FarmPlotID int64
	Persona    string
	ImageURL   string
	VideoURL   string
	CreatedAt  time.Time
}
