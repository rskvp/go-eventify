package models

import (
	"time"
)

const (
	FlowTypeDefault = "API"
	FlowTypeCron    = "CRON"
)

type Flow struct {
	ID             string    `json:"id"`
	Name           string    `json:"name" binding:"required" gorm:"unique"`
	Description    string    `json:"description"`
	Type           string    `json:"type" binding:"required"`
	Cron           string    `json:"cron"`
	Events         []*Event  `json:"events" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time `json:"createdAt"`
	LastExecAt     time.Time `json:"lastExecAt"`
	FlowExecutions []*FlowExecution
}

type FlowExecution struct {
	ID         string
	FlowID     string
	Input      string
	Output     string
	StartedAt  time.Time
	FinishedAt time.Time
}
