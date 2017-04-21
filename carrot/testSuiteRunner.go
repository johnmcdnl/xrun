package carrot

import (
	"fmt"
)

type TestSuiteRunner struct {
	suiteStepDefs []*StepDefinition
	TestFeatures  []*TestFeature `json:"features,omitempty"`
	MissingSteps  []*TestStep
}


func (tsr *TestSuiteRunner)Run() {
	for _, feature := range tsr.TestFeatures {
		testfeatureSync.Add(1)
		go tsr.RunTestFeature(feature)
	}
	testfeatureSync.Wait()
}

func (tsr *TestSuiteRunner)PrintMissingStepDefinitions() {
	fmt.Println("Missing Step Defs: ", tsr.MissingSteps)
}

func (tsr *TestSuiteRunner)BuildTestFeatures(dir string) {
	files, _ := filesWithExt(dir, ".feature")
	for _, featureFilePath := range files {
		var tf TestFeature
		tf.BuildTestFeatures(featureFilePath)
		tsr.TestFeatures = append(tsr.TestFeatures, &tf)
	}

}