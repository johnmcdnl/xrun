package carrot

import (
	"github.com/cucumber/gherkin-go"
)

type TestStep struct {
	*gherkin.PickleStep
}

func (tsr *TestSuiteRunner)RunTestStep(tc *TestStep) {

}

func (ts *TestStep)BuildTestStep(step *gherkin.PickleStep) {
	ts.PickleStep = step
}