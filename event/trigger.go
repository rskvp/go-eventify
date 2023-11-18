package event

type Trigger struct {
	id   string
	data string
}

func NewTrigger(id string, d string) *Trigger {
	return &Trigger{id, d}
}
