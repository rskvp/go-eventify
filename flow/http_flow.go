package flow

import (
	"log"

	"github.com/google/uuid"

	"assalielmehdi/eventify/event"
)

type HttpFlow struct {
	FlowCommon
}

func NewHttpFlow(e *event.Event) *HttpFlow {
	return &HttpFlow{
		FlowCommon: FlowCommon{
			Id:      uuid.NewString(),
			Starter: e,
		},
	}
}

func (f *HttpFlow) Register() {
	log.Printf("Registering manual flow=%s\n", f.Id)
}

func (f *HttpFlow) Run() {
	log.Printf("Running http flow=%s\n", f.Id)

	f.Starter.Run("")
}
