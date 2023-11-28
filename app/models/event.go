package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Action      *Action   `json:"-"`
	NextID      string    `json:"nextId"`
	PrevID      string    `json:"prevId"`
	FlowID      string    `json:"flowId" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	PositionX   float64
	PositionY   float64
	IsInput     bool
	IsOutput    bool
}

func NewEvent() *Event {
	return &Event{
		ID:     uuid.NewString(),
		Action: NewAction(),
	}
}
