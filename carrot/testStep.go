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

func (tsr *TestSuiteRunner)RunTestStep(ctx context.Context, testStep *TestStep) {
	defer testStepSync.Done()
	sd := tsr.findMatchingStepDefinition(testStep.Text)
	if sd == nil || sd.match == nil {
		tsr.AddMissingTestStep(testStep)
		return
	}
	err := sd.Execute(ctx, testStep)
	fmt.Sprint(err)

}

func (tsr *TestSuiteRunner)AddMissingTestStep(ts *TestStep) {
	tsr.MissingSteps = append(tsr.MissingSteps, ts)
}

func (ts *TestStep)BuildTestStep(step *gherkin.PickleStep) {
	ts.Id = uuid.NewV4().String()
	ts.PickleStep = step
}