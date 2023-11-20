package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          string
	Name        string
	Description string
	Action      *Action
	Next        *Event `gorm:"foreignKey:PrevEventID"`
	PrevEventID string
	FlowID      string
	CreatedAt   time.Time
}

func NewEvent() *Event {
	return &Event{
		ID:     uuid.NewString(),
		Action: NewAction(),
	}
}
