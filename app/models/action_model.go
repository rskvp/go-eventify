package models

const (
	ActionTypeIndentity = 1
)

type Action struct {
	ID      string
	Type    int
	EventID string
}

func NewAction() *Action {
	return &Action{}
}
