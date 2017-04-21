package carrot

import (
	"github.com/cucumber/gherkin-go"
	"github.com/satori/uuid"
	"context"
	"fmt"
)

type TestCase struct {
	Id        string
	*gherkin.Pickle
	TestSteps []*TestStep `json:"steps,omitempty"`
}


func (tsr *TestSuiteRunner)RunTestCase(tc *TestCase) {
	defer testCaseSync.Done()
	ctx := context.WithValue(context.Background(), "id", tc.Id)
	fmt.Println(ctx)
	for _, testStep := range tc.TestSteps {
		testStepSync.Add(1)
		go tsr.RunTestStep(ctx, testStep)
	}
	testStepSync.Wait()
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