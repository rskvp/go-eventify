package flow

import "assalielmehdi/eventify/event"

type FlowCommon struct {
	Id      string
	Starter *event.Event
}

type Flow interface {
	Register()
	Run()
}
