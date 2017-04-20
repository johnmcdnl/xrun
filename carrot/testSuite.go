package carrot

type TestSuite struct {
	TestSuiteRunner TestSuiteRunner
}

func (suite *TestSuite)Run() {
	suite.TestSuiteRunner.Run()
}

func (suite *TestSuite)Build() {
	var featureDir = "./internal"
	suite.TestSuiteRunner.BuildTestFeatures(featureDir)
	printJSON(suite)
}





