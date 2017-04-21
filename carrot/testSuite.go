package carrot

type TestSuite struct {
	TestSuiteRunner TestSuiteRunner
}

const featureDir = "./internal"

func (suite *TestSuite)Run() {
	suite.TestSuiteRunner.suiteStepDefs = GlobalStepDefinition

	suite.TestSuiteRunner.BuildTestFeatures(featureDir)

	suite.TestSuiteRunner.Run()

	suite.TestSuiteRunner.PrintMissingStepDefinitions()
}

func (suite *TestSuite)BuildAndRun() {
	imports, _ := WriteImportMarkers(featureDir)
	WriteMainTestFile(imports)
	RunMainTest()
}





