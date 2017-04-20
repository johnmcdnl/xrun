package main

import "github.com/johnmcdnl/xrun/carrot"

func main() {
//	xrun.BuildAndRunDir("./internal")
	//new(xrun.Runner).New().Run()
	new(carrot.TestSuite).Build()
	new(carrot.TestSuite).Run()
}
