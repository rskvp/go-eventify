package pubsub

type callback func(string)

type PubSub struct {
	subs map[string][]callback
}

func NewPubSub() *PubSub {
	return &PubSub{
		subs: make(map[string][]callback),
	}
}

func (ps *PubSub) Pub(id string, data string) *PubSub {
	subs, ok := ps.subs[id]

	if ok {
		for _, cb := range subs {
			cb(data)
		}
	}

	return ps
}

func (ps *PubSub) Sub(id string, cb callback) *PubSub {
	_, ok := ps.subs[id]

	if !ok {
		ps.subs[id] = make([]callback, 0)
	}

	ps.subs[id] = append(ps.subs[id], cb)

	return ps
}
