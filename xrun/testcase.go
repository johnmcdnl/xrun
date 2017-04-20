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

func (suite *Suite) runScenario(tCtx *TestContext, scenario *Scenario) {
	color.New(color.FgWhite, color.Bold).Print("\t Scenario: ")
	color.White(scenario.Name)

	var scenarioIsFailed bool
	for _, step := range scenario.Steps {
		//Make function call so can be go xxxFun()
		if scenarioIsFailed{
			color.Yellow(fmt.Sprint("\t\t ", step.Text))
			continue
		}
		suite.runStep(tCtx, step)
		if step.IsPassed{
			step.IsPassed = true
			color.Green(fmt.Sprint("\t\t ", step.Text))
			scenarioIsFailed = false
		}
		if !step.IsPassed{
			step.IsPassed = false
			color.Red(fmt.Sprint("\t\t ", step.Text))
			scenarioIsFailed = true
		}
	}
	fmt.Println()
}
