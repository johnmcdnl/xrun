package carrot

type TestSuiteRunner struct {
	contexts      []*TestContext
	suiteStepDefs []StepDefinition
	TestFeatures  []*TestFeature
}

func (tsr *TestSuiteRunner)Run() {
	for _, feature := range tsr.TestFeatures {
		tsr.RunTestFeature(feature)
	}
}

func (tsr *TestSuiteRunner)BuildTestFeatures(dir string) {
	files, _ := filesWithExt(dir, ".feature")
	for _, featureFilePath := range files{
		var tf TestFeature
		tf.BuildTestFeatures(featureFilePath)
		tsr.TestFeatures=append(tsr.TestFeatures, &tf)
	}



}