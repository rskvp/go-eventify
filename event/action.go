package event

import (
	"log"

	"github.com/google/uuid"
)

type Callback func(string) string

type Action struct {
	Id string
	cb Callback
}

var DefaultAction *Action = NewAction(func(s string) string {
	return s
})

func NewAction(cb Callback) *Action {
	return &Action{
		Id: uuid.NewString(),
		cb: cb,
	}
}

func (a *Action) Run(d string) string {
	log.Printf("Executing action=%s with input=%s\n", a.Id, d)

	d = a.cb(d)

	log.Printf("Execution of action=%s finished with output=%s\n", a.Id, d)

	return d
}
