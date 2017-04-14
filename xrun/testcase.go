package xrun

import "fmt"

type TestCase struct {
	Steps []*Step
}

func (s *Suite)RunTestCase(tc *TestCase) {

	for _, step := range tc.Steps {
		s.RunStep(step)
	}

	fmt.Println("(tc *TestCase)Run()")
}
