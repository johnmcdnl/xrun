package carrot

import (
	"github.com/cucumber/gherkin-go"
	"os"
	"github.com/satori/uuid"
)

type TestFeature struct {
	Id        string
	*gherkin.Feature
	TestCases []*TestCase `json:"testCases,omitempty"`
	Children  []interface{} `json:"children,omitempty"`
}

func (tsr *TestSuiteRunner)RunTestFeature(testFeature *TestFeature) {
	defer testfeatureSync.Done()
	for _, testCase := range testFeature.TestCases {
		testCaseSync.Add(1)
		go tsr.RunTestCase(testCase)
	}
	testCaseSync.Wait()
}

func (tf *TestFeature)BuildTestFeatures(path string) {
	tf.Id = uuid.NewV4().String()
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