package xrun

import (
	"github.com/cucumber/gherkin-go"
)

type Scenario struct {
	*gherkin.Pickle
	Steps []*Step  `json:"steps"`
}

func (suite *Suite) runScenario(scenario *Scenario) {
	for _, step := range scenario.Steps {
		suite.runStep(step)
	}

}
