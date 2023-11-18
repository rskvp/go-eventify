package event

import (
	"fmt"

	"github.com/google/uuid"

	ps "assalielmehdi/eventify/pubsub"
)

type Event struct {
	id      string
	ps      *ps.PubSub
	trigger *Trigger
	action  *Action
}

func NewEvent(ps *ps.PubSub, t *Trigger, a *Action) *Event {
	return &Event{uuid.New().String(), ps, t, a}
}

// Subscribes the event to its trigger
func (e *Event) Start() {
	e.ps.Sub(e.trigger.id, func(d string) {
		d = e.action.cb(d)

		e.ps.Pub(e.endId(), d)
	})
}

func (e *Event) startId() string {
	return fmt.Sprintf("START_%s", e.id)
}

func (e *Event) endId() string {
	return fmt.Sprintf("END_%s", e.id)
}
