package models

import (
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required" gorm:"unique"`
	Description string    `json:"description"`
	NextID      string    `json:"nextId"`
	PrevID      string    `json:"prevId"`
	FlowID      string    `json:"flowId" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	PositionX   float64
	PositionY   float64
	IsInput     bool
	IsOutput    bool
}

type EventExecution struct {
	ID         string
	EventID    string
	Input      string
	Output     string
	StartedAt  time.Time
	FinishedAt time.Time
}
