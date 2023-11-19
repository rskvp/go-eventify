package flow

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"

	"assalielmehdi/eventify/event"
)

type CronFlow struct {
	FlowCommon
	Cron string
}

func NewCronFlow(e *event.Event, c string) *CronFlow {
	return &CronFlow{
		FlowCommon: FlowCommon{
			Id:      uuid.NewString(),
			Starter: e,
		},
		Cron: c,
	}
}

func (f *CronFlow) Register() {
	log.Printf("Registering cron flow=%s with cron=%s\n", f.Id, f.Cron)

	s := gocron.NewScheduler(time.UTC)

	s.Cron(f.Cron).Do(func() {
		f.Run()
	})
}

func (f *CronFlow) Run() {
	log.Printf("Running http flow=%s\n", f.Id)

	f.Starter.Run("")
}
