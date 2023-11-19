package flow

import (
	"github.com/google/uuid"

	"assalielmehdi/eventify/event"
)

type HttpFlow struct {
	Flow
}

func NewHttpFlow(e *event.Event) *HttpFlow {
	return &HttpFlow{
		Flow: Flow{
			Id:      uuid.NewString(),
			Starter: e,
		},
	}
}
