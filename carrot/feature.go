package carrot

import "github.com/cucumber/gherkin-go"

type Feature struct {
	*gherkin.Feature
	Scenarios []Scenario
}

func (s *Suite)RunFeature(f *Feature) {
	for _, s := range f.Scenarios {
		s.RunScenario()
	}
}