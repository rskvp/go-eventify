package flow

import (
	"github.com/google/uuid"

	"assalielmehdi/eventify/event"
)

type ManualFlow struct {
	Flow
}

func NewManualFlow(e *event.Event) *ManualFlow {
	return &ManualFlow{
		Flow: Flow{
			Id:      uuid.NewString(),
			Starter: e,
		},
	}
}
