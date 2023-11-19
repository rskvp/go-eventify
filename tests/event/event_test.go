package event_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"assalielmehdi/eventify/event"
)

func TestEventRun(t *testing.T) {
	assert := assert.New(t)

	e1, e2, e3 := event.NewEvent(), event.NewEvent(), event.NewEvent()

	e1.Next = append(e1.Next, e2, e3)

	d1, d2, d3 := "", "", ""

	a1 := event.NewAction(func(d string) string {
		d1 = "d1"
		return d1
	})
	a2 := event.NewAction(func(d string) string {
		d2 = fmt.Sprintf("%s_d2", d)
		return d2
	})
	a3 := event.NewAction(func(d string) string {
		d3 = fmt.Sprintf("%s_d3", d)
		return d3
	})

	e1.Action = a1
	e2.Action = a2
	e3.Action = a3

	e1.Run("")

	assert.Equal("d1", d1)
	assert.Equal("d1_d2", d2)
	assert.Equal("d1_d3", d3)
}

func TestEventStartId(t *testing.T) {
	assert := assert.New(t)

	id := "id"
	e := event.NewEvent()
	e.Id = id

	startId := e.StartId()

	assert.Equal("START_id", startId)
}
