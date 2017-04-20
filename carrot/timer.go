package carrot

import "time"

type Timers struct {
	Timers []*Timer
}

type Timer struct {
	Id         string
	Name       string
	Executions []*Execution
}

type Execution struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
}
