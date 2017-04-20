package carrot

import (
	"github.com/cucumber/gherkin-go"
	"os"
)

type TestFeature struct {
	*gherkin.Feature
	TestCases []*TestCase `json:"testCases,omitempty"`
	Children  []interface{} `json:"children,omitempty"`
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
	tf.Feature = gd.Feature
	for _, pickle := range gd.Pickles() {
		var tc TestCase
		tc.BuildTestCase(pickle)
		tf.TestCases = append(tf.TestCases, &tc)
	}
}