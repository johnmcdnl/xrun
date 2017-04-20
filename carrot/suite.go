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

func (s *Suite)Run() {
	for _, f := range s.Features {
		f.RunFeature(f)
	}
}





