package domain

import "time"

type FarmPlot struct {
	ID          int64
	Name        string
	Location    string
	SizeSQM     int32
	MonthlyRent int32
	CropType    string
	Status      string
	CreatedAt   time.Time
}
