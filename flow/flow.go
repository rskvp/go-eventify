package flow

import (
	"assalielmehdi/eventify/event"
	"log"
)

type Flow struct {
	Id      string
	Starter *event.Event
}

func (f *Flow) Run(d string) {
	log.Printf("Running flow=%s\n", f.Id)

	f.Starter.Run(d)
}
