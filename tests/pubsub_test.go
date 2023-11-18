package tests

import (
	"testing"

	ps "assalielmehdi/eventify/pubsub"
)

func TestPubSub(t *testing.T) {
	ps := ps.NewPubSub()

	id, data, rcvData := "id", "data", ""

	ps.Sub(id, func(data string) {
		rcvData = data
	})

	ps.Pub(id, data)

	if rcvData != data {
		t.Fatalf("Given: %v, Expected %v", rcvData, data)
	}
}
