package xrun

import "fmt"

type TestCase struct {

}

func (tc *TestCase)Run() {
	fmt.Println("(tc *TestCase)Run()")
}
