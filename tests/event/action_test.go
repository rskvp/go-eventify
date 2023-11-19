package event_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"assalielmehdi/eventify/event"
)

func TestActionRun(t *testing.T) {
	assert := assert.New(t)

	cb := func(d string) string {
		return strings.ToLower(d)
	}
	a := event.NewAction(cb)

	d := a.Run("TEST")

	assert.Equal("test", d)
}
