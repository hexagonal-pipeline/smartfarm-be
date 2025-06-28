package raid

import (
	"smartfarm-be/internal/domain"
	"time"
)

type RaidResponse struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	CropType         string    `json:"crop_type"`
	TargetQuantity   int32     `json:"target_quantity"`
	MinParticipation int32     `json:"min_participation"`
	MaxParticipation int32     `json:"max_participation"`
	PricePerKg       int32     `json:"price_per_kg"`
	Deadline         time.Time `json:"deadline"`
	Status           string    `json:"status"`
	CreatorNickname  string    `json:"creator_nickname"`
	CreatedAt        time.Time `json:"created_at"`
	CurrentQuantity  int64     `json:"current_quantity"`
	ParticipantCount int64     `json:"participant_count"`
}

type RaidParticipationResponse struct {
	ID                  int64     `json:"id"`
	RaidID              int32     `json:"raid_id"`
	ParticipantNickname string    `json:"participant_nickname"`
	Quantity            int32     `json:"quantity"`
	ExpectedRevenue     int32     `json:"expected_revenue"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
}

type JoinRaidRequest struct {
	Nickname string `json:"nickname"`
	Quantity int32  `json:"quantity"`
}

func NewRaidResponse(raid domain.Raid) RaidResponse {
	return RaidResponse{
		ID:               raid.ID,
		Title:            raid.Title,
		Description:      raid.Description,
		CropType:         raid.CropType,
		TargetQuantity:   raid.TargetQuantity,
		MinParticipation: raid.MinParticipation,
		MaxParticipation: raid.MaxParticipation,
		PricePerKg:       raid.PricePerKg,
		Deadline:         raid.Deadline,
		Status:           raid.Status,
		CreatorNickname:  raid.CreatorNickname,
		CreatedAt:        raid.CreatedAt,
		CurrentQuantity:  raid.CurrentQuantity,
		ParticipantCount: raid.ParticipantCount,
	}
}

func NewRaidListResponse(raids []domain.Raid) []RaidResponse {
	responses := make([]RaidResponse, len(raids))
	for i, raid := range raids {
		responses[i] = NewRaidResponse(raid)
	}
	return responses
}

func NewRaidParticipationResponse(participation domain.RaidParticipation) RaidParticipationResponse {
	return RaidParticipationResponse{
		ID:                  participation.ID,
		RaidID:              participation.RaidID,
		ParticipantNickname: participation.ParticipantNickname,
		Quantity:            participation.Quantity,
		ExpectedRevenue:     participation.ExpectedRevenue,
		Status:              participation.Status,
		CreatedAt:           participation.CreatedAt,
	}
}

func NewRaidParticipationListResponse(participations []domain.RaidParticipation) []RaidParticipationResponse {
	responses := make([]RaidParticipationResponse, len(participations))
	for i, participation := range participations {
		responses[i] = NewRaidParticipationResponse(participation)
	}
	return responses
}
