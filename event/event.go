package event

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Event struct {
	Id     string
	Action *Action
	Next   []*Event
}

func NewEvent() *Event {
	return &Event{
		Id:     uuid.NewString(),
		Action: DefaultAction,
		Next:   make([]*Event, 0),
	}
}

func (e *Event) Run(d string) {
	log.Printf("Running event=%s with input=%s\n", e.Id, d)

	d = e.Action.Run(d)

	log.Printf("Excution of event=%s finished with output=%s\n", e.Id, d)

	for _, next := range e.Next {
		next.Run(d)
	}
}

func (e *Event) StartId() string {
	return fmt.Sprintf("START_%s", e.Id)
}
