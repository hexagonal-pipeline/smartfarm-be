package domain

import (
	"time"
)

type FarmPlot struct {
	ID            int64
	Name          string
	Location      string
	SizeSqm       int32
	MonthlyRent   int32
	CropType      string
	Status        string
	CreatedAt     time.Time
	PersonaPrompt *string
}

type Rental struct {
	// ... existing code ...
}
