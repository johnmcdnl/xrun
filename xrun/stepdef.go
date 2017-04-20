package xrun

import (
	"regexp"
	"reflect"
	"fmt"
	"strconv"
)

type StepDef struct {
	Regexp  *regexp.Regexp
	Handler interface{}
	Args    []string
}

func (sd *StepDef) runStepDef(tCtx *TestContext) []*TestError {
	if sd == nil {
		//TODO should return (found bool, err error)
		var errs []*TestError
		errs = append(errs, &TestError{})
		return errs
	}

	if len(sd.Args) == 0 {

	}

	hv := reflect.ValueOf(sd.Handler)
	hv.Call(sd.getValues(tCtx))

	return tCtx.T().TestErrors
}

func (sd *StepDef)getValues(tCtx *TestContext) []reflect.Value {
	fn := sd.Handler
	args := sd.Args
	fnType := reflect.ValueOf(fn).Type()

	var values []reflect.Value
	values = append(values, reflect.ValueOf(tCtx))

	for m, i := 0, 0; i < fnType.NumIn(); i++ {
		param := fnType.In(i)
		var v interface{}
		switch param.Kind() {
		case reflect.TypeOf(&TestContext{}).Kind():
			continue
		case reflect.String:
			v = args[m]
			m++
		case reflect.Int:
			i, _ := strconv.ParseInt(args[m], 10, 32)
			v = int(i)
			m++
		case reflect.Int64:
			i, _ := strconv.ParseInt(args[m], 10, 32)
			v = int64(i)
			m++
		case reflect.Float64:
			v, _ = strconv.ParseFloat(args[m], 64)
			m++
		default:
			fmt.Println(args)
		}

		values = append(values, reflect.ValueOf(v))
	}

	return values
}