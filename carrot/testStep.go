package carrot

import (
	"github.com/cucumber/gherkin-go"
	"fmt"
)

type TestStep struct {
	*gherkin.PickleStep
}

func (tsr *TestSuiteRunner)RunTestStep(tc *TestStep) {
	fmt.Println(tc.Text)
}

func (ts *TestStep)BuildTestStep(step *gherkin.PickleStep) {
	ts.PickleStep = step
}