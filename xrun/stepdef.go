package xrun

import (
	"regexp"
	"reflect"
)

type StepDef struct {
	Regexp  *regexp.Regexp
	Handler interface{}
	Args    []interface{}
}

func (sd *StepDef) Run() {
	var values []reflect.Value
	for _, arg := range sd.Args {
		values = append(values, reflect.ValueOf(arg))
	}
	hv := reflect.ValueOf(sd.Handler)
	hv.Call(values)
}

