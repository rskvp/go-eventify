package event

type Callback func(string) string

type Action struct {
	cb Callback
}

func NewAction(cb Callback) *Action {
	return &Action{cb}
}
