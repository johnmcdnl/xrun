package xrun

import (
	"regexp"
)

//Could create global Suite and clone it when execution starts
var GlobalStepDefStore []*StepDef

func store(match string, fn interface{}) {
	var stepDef StepDef
	stepDef.Regexp = regexp.MustCompile(match)
	stepDef.Handler = fn
	GlobalStepDefStore = append(GlobalStepDefStore, &stepDef)
}

func BeforeAll(){}
func Before(...interface{}){}
func After(...interface{}){}
func AfterAll(...interface{}){}

func Given(match string, fn interface{}) {
	store(match, fn)
}
func When(match string, fn interface{}) {
	store(match, fn)
}
func Then(match string, fn interface{}) {
	store(match, fn)
}
func And(match string, fn interface{}) {
	store(match, fn)
}
func But(match string, fn interface{}) {
	store(match, fn)
}
