package carrot

import (
	"regexp"
	"fmt"
	"reflect"
	"strconv"
	"context"
)

var GlobalStepDefinition []*StepDefinition

func Step(text string, fn interface{}) {
	var sd = StepDefinition{match: regexp.MustCompile(text), fn:fn}
	GlobalStepDefinition = append(GlobalStepDefinition, &sd)
}

type StepDefinition struct {
	match  *regexp.Regexp
	fn     interface{}
	params []string //TODO I'd much rather to be []interface{}
}

func (tsr *TestSuiteRunner)findMatchingStepDefinition(text string) *StepDefinition {
	var matchSd *StepDefinition
	for _, sd := range tsr.suiteStepDefs {
		params, isMatch := sd.isMatch(text)
		if !isMatch {
			continue
		}
		if matchSd == nil {
			matchSd = sd
			if len(params) > 0 {
				matchSd.params = params[1:]
			}
			continue
		}
		if len(sd.match.String()) > len(sd.match.String()) {
			matchSd = sd
			if len(params) > 0 {
				matchSd.params = params[1:]
			}
			continue
		}

	}
	return matchSd
}

func (sd *StepDefinition)isMatch(text string) ([]string, bool) {
	matches := sd.match.FindStringSubmatch(text);
	if matches == nil {
		return matches, false
	}
	return matches, true
}

func (sd *StepDefinition)Execute(tCtx context.Context, step *TestStep) error {

	fmt.Println(sd.params, fmt.Sprint(step.Arguments))
	if len(step.Arguments) > 0 {
		fmt.Println(reflect.TypeOf(step.Arguments[0]))
	}

	var values []reflect.Value
	values = append(values, reflect.ValueOf(tCtx))

	fnType := reflect.ValueOf(sd.fn).Type()
	for m, i := 0, 0; i < fnType.NumIn(); i++ {
		param := fnType.In(i)
		var v interface{}
		switch param.Kind() {
		case reflect.TypeOf(&TestContext{}).Kind():
			continue
		case reflect.String:
			v = sd.params[m]
			m++
		case reflect.Int:
			i, _ := strconv.ParseInt(sd.params[m], 10, 32)
			v = int(i)
			m++
		case reflect.Int64:
			i, _ := strconv.ParseInt(sd.params[m], 10, 32)
			v = int64(i)
			m++
		case reflect.Float64:
			v, _ = strconv.ParseFloat(sd.params[m], 64)
			m++
		default:
			fmt.Println(sd.params)
		}

		values = append(values, reflect.ValueOf(v))
	}

	fmt.Println(tCtx.Id)
	return nil
}