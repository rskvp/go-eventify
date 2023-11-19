package runners

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"

	"assalielmehdi/eventify/flow"
)

type CronRunner struct{}

func NewCronRunner() *CronRunner {
	return &CronRunner{}
}

func (r *CronRunner) Schedule(f *flow.CronFlow, d string) {
	log.Printf("Scheduling cron flow=%s with cron=%s\n", f.Id, f.Cron)

	s := gocron.NewScheduler(time.UTC)

	s.Cron(f.Cron).Do(func() {
		f.Run(d)
	})
}
