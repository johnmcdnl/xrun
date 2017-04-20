package carrot

import "github.com/cucumber/gherkin-go"

type Feature struct {
	*gherkin.Feature
	Scenarios []Scenario
}

func (suite *Suite)RunFeature(f *Feature) {
	for _, scenario := range f.Scenarios {
		suite.RunScenario(scenario)
	}
}