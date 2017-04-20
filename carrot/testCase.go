package carrot

import (
	"github.com/cucumber/gherkin-go"
)

type TestCase struct {
	*gherkin.Pickle
	Steps []*TestStep
}

func (tsr *TestSuiteRunner)RunTestCase(tc *TestCase) {
	for _, ts := range tc.Steps {
		tsr.RunTestStep(ts)
	}
}

func (tc *TestCase)BuildTestCase(pickle *gherkin.Pickle) {
	for _, step := range pickle.Steps {
		var ts TestStep
		ts.BuildTestStep(step)
		tc.Steps=append(tc.Steps, &ts)
	}
}