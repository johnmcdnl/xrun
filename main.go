package main

import (
	"fmt"
	"github.com/johnmcdnl/xrun/xrun"
)

func main() {
	fmt.Println("XRUN v0.1")
	var r xrun.Runner
	var tc xrun.TestCase
	r.Suite.TestCases = append(r.Suite.TestCases, tc)
	r.Suite.TestCases = append(r.Suite.TestCases, tc)
	r.Suite.TestCases = append(r.Suite.TestCases, tc)
	r.Run()
}

