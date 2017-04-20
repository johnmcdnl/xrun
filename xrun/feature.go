package xrun

import (
	"github.com/cucumber/gherkin-go"
	"github.com/fatih/color"
	"github.com/satori/uuid"
)

type Feature struct {
	*gherkin.Feature
	Scenarios []*Scenario  `json:"scenarios"`
	Children interface{} `json:"children,omitempty"`
}


func (s *Suite)runFeature(f *Feature) {
	color.New(color.FgWhite, color.Bold).Print("Feature: ")
	color.White(f.Name)
	for _, scenario := range f.Scenarios {
		tCtx := s.getContext(uuid.NewV4().String())
		s.runScenario(tCtx, scenario)
	}
}