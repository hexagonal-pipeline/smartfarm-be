package domain

type FarmPlot struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	IsAvailable bool   `json:"isAvailable"`
}
