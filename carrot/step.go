package carrot

import "github.com/cucumber/gherkin-go"

type Step struct {
	*gherkin.PickleStep
	Result Result
}

func (suite *Suite)RunStep(step *Step) {

}
