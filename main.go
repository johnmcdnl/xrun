package main

import "github.com/johnmcdnl/xrun/carrot"

func main() {
//	xrun.BuildAndRunDir("./internal")
	//new(xrun.Runner).New().Run()
	var ts carrot.TestSuite
	ts.Build()
	ts.Run()
}
