package carrot

import (
	"regexp"
	"fmt"
)

type StepDefinition struct {
	match  *regexp.Regexp
	fn     interface{}
	params []string //TODO I'd much rather to be []interface{}
}

func (tsr *TestSuiteRunner)findMatchingStepDefinition(text string) *StepDefinition {
	var matchSd *StepDefinition
	for _, sd := range tsr.suiteStepDefs {
		params, isMatch := sd.isMatch(text)
		if !isMatch {
			continue
		}
		if matchSd == nil {
			matchSd = sd
			if len(params) > 0 {
				matchSd.params = params[1:]
			}
			continue
		}
		if len(sd.match.String()) > len(sd.match.String()) {
			matchSd = sd
			if len(params) > 0 {
				matchSd.params = params[1:]
			}
			continue
		}

	}
	fmt.Println(matchSd)
	return matchSd
}

func (sd *StepDefinition)isMatch(text string) ([]string, bool) {
	matches := sd.match.FindStringSubmatch(text);
	if matches == nil {
		return matches, false
	}
	return matches, true
}