package carrot

import (
	"fmt"
)

type TestSuiteRunner struct {
	contexts      []*TestContext
	suiteStepDefs []*StepDefinition
	TestFeatures  []*TestFeature `json:"features,omitempty"`
	MissingSteps  []*TestStep
}

func (tsr *TestSuiteRunner)Run() {
	for _, feature := range tsr.TestFeatures {
		tsr.RunTestFeature(feature)
	}
}

func (tsr *TestSuiteRunner)GetContext(id string) *TestContext {
	for _, tCtx := range tsr.contexts {
		if tCtx.Id == id {
			return tCtx
		}
	}
	tCtx := NewContext(id)
	tsr.contexts = append(tsr.contexts, tCtx)
	return tsr.GetContext(id)
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