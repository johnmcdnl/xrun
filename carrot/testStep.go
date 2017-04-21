package carrot

import (
	"github.com/cucumber/gherkin-go"
	"fmt"
	"github.com/satori/uuid"
	"context"
)

type TestStep struct {
	Id string
	*gherkin.PickleStep
}

func (tsr *TestSuiteRunner)RunTestStep(tCtx context.Context, ts *TestStep) {
	sd := tsr.findMatchingStepDefinition(ts.Text)
	if sd == nil || sd.match == nil {
		tsr.AddMissingTestStep(ts)
		return
	}
	err := sd.Execute(tCtx, ts)
	fmt.Sprint(err)
}

func (tsr *TestSuiteRunner)AddMissingTestStep(ts *TestStep) {
	tsr.MissingSteps = append(tsr.MissingSteps, ts)
}

func (ts *TestStep)BuildTestStep(step *gherkin.PickleStep) {
	ts.Id = uuid.NewV4().String()
	ts.PickleStep = step
}