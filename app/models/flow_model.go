package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	FlowTypeDefault = 1
	FlowTypeCron    = 2
)

type Flow struct {
	ID          string
	Name        string
	Description string
	Type        int
	Cron        string
	Events      []Event
	CreatedAt   time.Time
	LastExecAt  time.Time
}

func NewFlow() *Flow {
	return &Flow{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Type:      FlowTypeDefault,
	}
}
