package runners

import (
	"log"

	"assalielmehdi/eventify/flow"
)

type HttpRunner struct {
	Flows map[string]*flow.HttpFlow
}

func NewHttpRunner() *HttpRunner {
	return &HttpRunner{
		Flows: make(map[string]*flow.HttpFlow),
	}
}

func (r *HttpRunner) Register(f *flow.HttpFlow) {
	log.Printf("Registering http flow=%s\n", f.Id)

	r.Flows[f.Id] = f
}

func (r *HttpRunner) Run(id string) {
	f, ok := r.Flows[id]
	if !ok {
		log.Printf("No registered http flow=%s\n", id)
		return
	}

	f.Run("")
}
