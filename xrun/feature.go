package xrun

import (
	"fmt"
	"github.com/cucumber/gherkin-go"
)

type Feature struct {
	*gherkin.Feature
	Scenarios []*Scenario  `json:"scenarios"`
	Children interface{} `json:"children,omitempty"`
}


func (s *Suite)RunFeature(f *Feature) {
	for _, scenario := range f.Scenarios {
		s.RunScenario(scenario)
	}

	fmt.Println("(tc *TestCase)Run()")
}