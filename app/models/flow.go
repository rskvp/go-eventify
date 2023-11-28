package models

import (
	"time"
)

const (
	FlowTypeDefault = "API"
	FlowTypeCron    = "CRON"
)

type Flow struct {
	ID          string    `json:"id" `
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Type        string    `json:"type" binding:"required"`
	Cron        string    `json:"cron"`
	Events      []*Event  `json:"events"`
	CreatedAt   time.Time `json:"createdAt"`
	LastExecAt  time.Time `json:"lastExecAt"`
}
