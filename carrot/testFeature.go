package carrot

import (
	"github.com/cucumber/gherkin-go"
	"os"
)

type TestFeature struct {
	TestCases []*TestCase
}

func (tsr *TestSuiteRunner)RunTestFeature(testFeature *TestFeature) {
	for _, testCase := range testFeature.TestCases {
		tsr.RunTestCase(testCase)
	}
}

func (tf *TestFeature)BuildTestFeatures(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	gd, _ := gherkin.ParseGherkinDocument(file)
	for _, pickle := range gd.Pickles() {
		var tc TestCase
		tc.BuildTestCase(pickle)
		tf.TestCases = append(tf.TestCases, &tc)
	}
}