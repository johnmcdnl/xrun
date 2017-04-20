package carrot

func NewSuite() *Suite {
	var s Suite

	return &s
}

type Suite struct {
	Contexts      []*TestContext
	SuiteStepDefs []StepDefinition
	Features      []*Feature
}

func (suite *Suite)Run() {
	for _, feature := range suite.Features {
		suite.RunFeature(feature)
	}
}





