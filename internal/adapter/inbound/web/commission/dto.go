package commission

import (
	"smartfarm-be/internal/domain"
	"time"
)

type CreateCommissionWorkRequest struct {
	Nickname        string `json:"nickname"`
	PlotID          int32  `json:"plot_id"`
	TaskType        string `json:"task_type"`
	TaskDescription string `json:"task_description"`
	CreditCost      int32  `json:"credit_cost"`
}

type CommissionWorkResponse struct {
	ID                int64      `json:"id"`
	RequesterNickname string     `json:"requester_nickname"`
	PlotID            int32      `json:"plot_id"`
	TaskType          string     `json:"task_type"`
	TaskDescription   string     `json:"task_description"`
	Status            string     `json:"status"`
	CreditCost        int32      `json:"credit_cost"`
	RequestedAt       time.Time  `json:"requested_at"`
	CompletedAt       *time.Time `json:"completed_at,omitempty"`
}

type CommissionWorkListResponse struct {
	Items []*CommissionWorkResponse `json:"items"`
}

func NewCommissionWorkResponse(cw *domain.CommissionWork) *CommissionWorkResponse {
	return &CommissionWorkResponse{
		ID:                cw.ID,
		RequesterNickname: cw.RequesterNickname,
		PlotID:            cw.PlotID,
		TaskType:          cw.TaskType,
		TaskDescription:   cw.TaskDescription,
		Status:            cw.Status,
		CreditCost:        cw.CreditCost,
		RequestedAt:       cw.RequestedAt,
		CompletedAt:       cw.CompletedAt,
	}
}

func NewCommissionWorkListResponse(works []domain.CommissionWork) *CommissionWorkListResponse {
	items := make([]*CommissionWorkResponse, len(works))
	for i, w := range works {
		items[i] = NewCommissionWorkResponse(&w)
	}
	return &CommissionWorkListResponse{Items: items}
}
