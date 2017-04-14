package xrun

import (
	"regexp"
	"reflect"
)

type StepDef struct {
	Regexp  *regexp.Regexp
	Handler interface{}
	Args    []string
}

func (sd *StepDef) Run() {
	if sd == nil {
		return
	}

	var values []reflect.Value
	for _, arg := range sd.Args {
		values = append(values, reflect.ValueOf(arg))
	}
	hv := reflect.ValueOf(sd.Handler)
	hv.Call(values)
}

