package xrun

import (
	"github.com/cucumber/gherkin-go"
	"github.com/pkg/errors"
	"fmt"
	"github.com/fatih/color"
)

type Step struct {
	*gherkin.PickleStep
}

func (s *Suite)runStep(step *Step) {
	stepDef, err := s.findMatchingStepDef(step)
	if err != nil {
		//fmt.Println(err)
	}
	stepDef.runStepDef()

	color.Green(fmt.Sprint("\t\t ", step.Text))
}

func (s *Suite)findMatchingStepDef(step *Step) (*StepDef, error) {
	var bestMatch *StepDef
	for _, sd := range s.StepDefs {
		if match := sd.Regexp.FindStringSubmatch(step.Text); match != nil {
			if bestMatch==nil{
				bestMatch=sd
			}
			if bestMatch.Regexp == nil || len(sd.Regexp.String()) > len(bestMatch.Regexp.String()) {
				bestMatch = sd
				bestMatch.Args = match[1:]
			}
		}
	}
	if bestMatch==nil{
		return nil, errors.New("No match found")
	}
	return bestMatch, nil
}