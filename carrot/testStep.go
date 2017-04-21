package carrot

import (
	"github.com/cucumber/gherkin-go"
)

type TestStep struct {
	*gherkin.PickleStep
}

func (tsr *TestSuiteRunner)RunTestStep(ts *TestStep) {
	match := tsr.findMatchingStepDefinition(ts.Text)
	if match == nil || match.match == nil {
		tsr.AddMissingTestStep(ts)
		return
	}
}

func (tsr *TestSuiteRunner)AddMissingTestStep(ts *TestStep) {
	tsr.MissingSteps = append(tsr.MissingSteps, ts)
}

func (ts *TestStep)BuildTestStep(step *gherkin.PickleStep) {
	ts.PickleStep = step
}