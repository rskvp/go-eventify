package flow

import (
	"log"

	"github.com/google/uuid"

	"assalielmehdi/eventify/event"
)

type ManualFlow struct {
	FlowCommon
}

func NewManualFlow(e *event.Event) *ManualFlow {
	return &ManualFlow{
		FlowCommon: FlowCommon{
			Id:      uuid.NewString(),
			Starter: e,
		},
	}
}

func (f *ManualFlow) Register() {
	log.Printf("Registering manual flow=%s\n", f.Id)
}

func (f *ManualFlow) Run() {
	log.Printf("Running manual flow=%s\n", f.Id)

	f.Starter.Run("")
}
