package carrot

import (
	"github.com/cucumber/gherkin-go"
	"github.com/satori/uuid"
)

type TestCase struct {
	Id        string
	*gherkin.Pickle
	TestSteps []*TestStep `json:"steps,omitempty"`
}

func (tsr *TestSuiteRunner)RunTestCase(tc *TestCase) {
	for _, ts := range tc.TestSteps {
		tsr.RunTestStep(tsr.GetContext(tc.Id), ts)
	}
}

func (tc *TestCase)BuildTestCase(pickle *gherkin.Pickle) {
	tc.Id = uuid.NewV4().String()
	tc.Pickle = pickle
	for _, step := range pickle.Steps {
		var ts TestStep
		ts.BuildTestStep(step)
		tc.TestSteps = append(tc.TestSteps, &ts)
	}
}