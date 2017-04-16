package xrun

import (
	"github.com/cucumber/gherkin-go"
	"github.com/fatih/color"
	"fmt"
)

type Scenario struct {
	*gherkin.Pickle
	Steps []*Step  `json:"steps"`
}

func (suite *Suite) runScenario(scenario *Scenario) {
	color.New(color.FgWhite, color.Bold).Print("\t Scenario: ")
	color.White(scenario.Name)
	fmt.Println()
	for _, step := range scenario.Steps {
		suite.runStep(step)
	}

}
