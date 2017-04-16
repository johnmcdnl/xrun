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


func (s *Suite)runFeature(f *Feature) {
	for _, scenario := range f.Scenarios {
		s.runScenario(scenario)
	}

	fmt.Println("(tc *TestCase)Run()")
}