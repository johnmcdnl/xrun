package xrun

import (
	"github.com/cucumber/gherkin-go"
	"fmt"
	"github.com/pkg/errors"
)

type Step struct {
	*gherkin.PickleStep
}

func (s *Suite)RunStep(step *Step) {
	fmt.Println(step.Text)
	stepDef, err := s.FindMatchingStepDef(step)
	if err != nil {
		//fmt.Println(err)
	}
	stepDef.Run()
}

func (s *Suite)FindMatchingStepDef(step *Step) (*StepDef, error) {
	var bestMatch *StepDef
	for _, sd := range s.StepDefs {
		if match := sd.Regexp.FindStringSubmatch(step.Text); match != nil {
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