package xrun

import (
	"fmt"
)

type Suite struct {
	TestCases []TestCase
	StepDefs []StepDef
}

func (s *Suite)Run() {
	fmt.Println("(s *Suite)Run()")
	for _, tc := range s.TestCases {
		tc.Run()
	}

	var sd StepDef
	var fn = func(){
		fmt.Println("hello")
	}
	sd.Handler=fn
	sd.Run()
}