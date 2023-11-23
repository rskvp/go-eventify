package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Action      *Action   `json:"-"`
	Next        *Event    `json:"next" gorm:"foreignKey:PrevEventID"`
	PrevEventID string    `json:"prevEventId"`
	FlowID      string    `json:"flowId" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewEvent() *Event {
	return &Event{
		ID:     uuid.NewString(),
		Action: NewAction(),
	}
}
