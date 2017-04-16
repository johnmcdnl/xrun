package xrun

type Suite struct {
	Features []*Feature `json:"features"`
	StepDefs  []*StepDef `json:"-"`
}

func newSuite() *Suite {
	var s Suite
	s.Features = make([]*Feature, 0)
	s.StepDefs = make([]*StepDef, 0)
	return &s
}

func (s *Suite)runSuite() {
	for _, f := range s.Features {
		s.runFeature(f)
	}
}