package carrot

import "github.com/cucumber/gherkin-go"

type Scenario struct {
	*gherkin.Pickle
	Steps []*Step  `json:"steps"`
}

func (suite *Suite)RunScenario(scenario *Scenario) {
	for _, step := range scenario.Steps {
		suite.RunStep(step)
	}
}
