package farm

import (
	"smartfarm-be/internal/domain"
	"time"
)

// FarmPlotResponse is the DTO for a single farm plot.
type FarmPlotResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Location      string `json:"location"`
	SizeSqm       int32  `json:"size_sqm"`
	MonthlyRent   int32  `json:"monthly_rent"`
	CropType      string `json:"crop_type"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	PersonaPrompt string `json:"persona_prompt"`
}

// NewFarmPlotResponse creates a new FarmPlotResponse from a domain model.
func NewFarmPlotResponse(plot domain.FarmPlot) FarmPlotResponse {
	return FarmPlotResponse{
		ID:          plot.ID,
		Name:        plot.Name,
		Location:    plot.Location,
		SizeSqm:     plot.SizeSqm,
		MonthlyRent: plot.MonthlyRent,
		CropType:    plot.CropType,
		Status:      plot.Status,
		CreatedAt:   plot.CreatedAt.Format(time.RFC3339),
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
