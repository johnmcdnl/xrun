package carrot

import (
	"github.com/cucumber/gherkin-go"
)

type TestCase struct {
	*gherkin.Pickle
	TestSteps []*TestStep `json:"steps,omitempty"`
}

func (tsr *TestSuiteRunner)RunTestCase(tc *TestCase) {
	for _, ts := range tc.TestSteps {
		tsr.RunTestStep(ts)
	}
}

func (tc *TestCase)BuildTestCase(pickle *gherkin.Pickle) {
	tc.Pickle = pickle
	for _, step := range pickle.Steps {
		var ts TestStep
		ts.BuildTestStep(step)
		tc.TestSteps =append(tc.TestSteps, &ts)
	}
}