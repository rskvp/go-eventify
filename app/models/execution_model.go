package models

import "time"

const (
	ExecutionTypeFlow  = 1
	ExecutionTypeEvent = 2
)

type Execution struct {
	ID         string
	Type       int
	ResourceId string
	Input      string
	Output     string
	StartedAt  time.Time
}
