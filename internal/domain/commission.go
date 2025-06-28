package domain

import "time"

type CommissionWork struct {
	ID                int64
	RequesterNickname string
	PlotID            int32
	TaskType          string
	TaskDescription   string
	Status            string
	CreditCost        int32
	RequestedAt       time.Time
	CompletedAt       *time.Time
}
