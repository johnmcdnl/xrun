package xrun

type Suite struct {
	Contexts []*TestContext `json:"-"`
	Features []*Feature `json:"features"`
	StepDefs []*StepDef `json:"-"`
}

func (s *Suite)getContext(id string) *TestContext {
	var tc *TestContext
	for _, ctx := range s.Contexts {
		if ctx.ID == id {
			tc = ctx
			break
		}
	}

	if tc == nil {
		tc = testContextWithId(id)
		s.Contexts = append(s.Contexts, tc)
	}

	return tc
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