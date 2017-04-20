package xrun

import (
	"github.com/cucumber/gherkin-go"
	"github.com/pkg/errors"
)

type Step struct {
	*gherkin.PickleStep
	StepResult
}

type StepResult struct {
	IsExecuted      bool `json:"isExecuted,omitempty"`
	IsPassed        bool `json:"isPassed,omitempty"`
	HasMatchingStep bool `json:"hasMatchingStep,omitempty"`
	Errors          []*TestError `json:"errors,omitempty"`
}

func (s *Suite)runStep(tCtx *TestContext, step *Step) {
	stepDef, err := s.findMatchingStepDef(tCtx, step)
	if err != nil {
		return
	}
	step.HasMatchingStep = true

	if testErrors := stepDef.runStepDef(tCtx); len(testErrors) > 0 {
		step.Errors = testErrors
		step.IsExecuted = true
		step.IsPassed = false
		return
	}
	step.IsExecuted = true
	step.IsPassed = true

}

func (s *Suite)findMatchingStepDef(tCtx *TestContext, step *Step) (*StepDef, error) {
	var bestMatch *StepDef
	for _, sd := range s.StepDefs {
		if match := sd.Regexp.FindStringSubmatch(step.Text); match != nil {
			if bestMatch == nil {
				bestMatch = sd
			}
			if bestMatch.Regexp != nil || len(sd.Regexp.String()) > len(bestMatch.Regexp.String()) {
				bestMatch = sd
				bestMatch.Args = match[1:]
			}
		}
	}
	if bestMatch == nil {
		return nil, errors.New("No match found")
	}

	return bestMatch, nil
}
