package flow

import (
	"github.com/google/uuid"

	"assalielmehdi/eventify/event"
)

type CronFlow struct {
	Flow
	Cron string
}

func NewCronFlow(e *event.Event, c string) *CronFlow {
	return &CronFlow{
		Flow: Flow{
			Id:      uuid.NewString(),
			Starter: e,
		},
		Cron: c,
	}
}
