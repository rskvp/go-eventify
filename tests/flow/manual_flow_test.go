package flow_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"assalielmehdi/eventify/event"
	"assalielmehdi/eventify/flow"
)

func TestManualFlowRegis(t *testing.T) {
	assert := assert.New(t)

	ok := false
	e := event.NewEvent()
	e.Action = event.NewAction(func(d string) string {
		ok = true
		return d
	})
	var f flow.Flow = flow.NewManualFlow(e)

	f.Run()

	assert.True(ok)
}
