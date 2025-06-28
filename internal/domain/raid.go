package domain

import "time"

type Raid struct {
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

type RaidParticipation struct {
	ID                  int64     `json:"id"`
	RaidID              int32     `json:"raid_id"`
	ParticipantNickname string    `json:"participant_nickname"`
	Quantity            int32     `json:"quantity"`
	ExpectedRevenue     int32     `json:"expected_revenue"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
}
