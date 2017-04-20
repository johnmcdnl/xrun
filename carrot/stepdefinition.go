package carrot

import "regexp"

type StepDefinition struct {
	match  regexp.Regexp
	fn     interface{}
	params []interface{}
}