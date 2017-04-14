package xrun

import "fmt"

type Suite struct {
	TestCases []TestCase
}

func (s *Suite)Run() {
	fmt.Println("(s *Suite)Run()")
	for _, tc := range s.TestCases {
		tc.Run()
	}
}