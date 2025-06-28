package plantcard

type CreatePlantCardRequest struct {
	FarmPlotID int64  `json:"farm_plot_id"`
	Event      string `json:"event"`
}

type CreatePlantCardResponse struct {
	ID           int64  `json:"id"`
	FarmPlotID   int64  `json:"farm_plot_id"`
	Persona      string `json:"persona"`
	ImageURL     string `json:"image_url"`
	VideoURL     string `json:"video_url"`
	EventMessage string `json:"event_message"`
	CreatedAt    string `json:"created_at"`
}
