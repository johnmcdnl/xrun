package xrun

import (
	"github.com/cucumber/gherkin-go"
	"fmt"
	"github.com/pkg/errors"
)

type Step struct {
	*gherkin.PickleStep
}

func (s *Suite)runStep(step *Step) {
	fmt.Println(step.Text)
	stepDef, err := s.findMatchingStepDef(step)
	if err != nil {
		//fmt.Println(err)
	}
	stepDef.runStepDef()
}

func (s *Suite)findMatchingStepDef(step *Step) (*StepDef, error) {
	var bestMatch *StepDef
	for _, sd := range s.StepDefs {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("sd", sd)
		fmt.Println()
		fmt.Println()
		fmt.Println()
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