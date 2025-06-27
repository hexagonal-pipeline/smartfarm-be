package farm

import "smartfarm-be/internal/domain"

// FarmPlotResponse is the DTO for a single farm plot.
type FarmPlotResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// NewFarmPlotResponse creates a new FarmPlotResponse from a domain model.
func NewFarmPlotResponse(plot domain.FarmPlot) FarmPlotResponse {
	return FarmPlotResponse{
		ID:   plot.ID,
		Name: plot.Name,
	}
}

// NewFarmPlotListResponse creates a list of FarmPlotResponse from a list of domain models.
func NewFarmPlotListResponse(plots []domain.FarmPlot) []FarmPlotResponse {
	res := make([]FarmPlotResponse, len(plots))
	for i, plot := range plots {
		res[i] = NewFarmPlotResponse(plot)
	}
	return res
}
