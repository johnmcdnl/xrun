package xrun

import (
	"github.com/pkg/errors"
	"runtime/debug"
)

type TestingT struct {
	TestErrors []*TestError
}

type TestError struct {
	error   error
	Message string `json:"errorMessage,omitempty"`
	Stack   string `json:"stackTrace,omitempty"`
}

func (t *TestingT)Errorf(format string, args ...interface{}) {
	err := errors.Errorf(format, args)
	var te = TestError{
		error: err,
		Stack:string(debug.Stack()),
		Message : err.Error(),
	}
	t.TestErrors = append(t.TestErrors, &te)
}