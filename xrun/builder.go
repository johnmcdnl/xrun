package xrun

import (
	"path/filepath"
	"os"
	"strings"
	"github.com/cucumber/gherkin-go"
	"io/ioutil"
	"encoding/json"
	"path"
)

var (
	stepDefRootDir = filepath.ToSlash(path.Join("internal"))
	featuresRootDir = filepath.ToSlash(path.Join("internal", "features"))
)

func (r *Runner)BuildAndRun() {
	os.MkdirAll(featuresRootDir, os.ModePerm)
	os.MkdirAll(stepDefRootDir, os.ModePerm)
	s := newSuite()
	s.Features = gherkinToFeatures()
	s.StepDefs = getStepDefs()
	r.Suite = s

	r.Suite.runSuite()
	r.Reporter.run()

	j, _ := json.MarshalIndent(r.Suite, "", "\t")
	ioutil.WriteFile("suite.json", j, os.ModePerm)
}

func getStepDefs() []*StepDef {
	return GlobalStepDefStore
}

func gherkinToFeatures() []*Feature {
	featureFiles, _ := filesWithExt(featuresRootDir, ".feature")

	var gherkinDocuments = []*gherkin.GherkinDocument{}

	for _, f := range featureFiles {
		file, _ := os.Open(f)
		gd, _ := gherkin.ParseGherkinDocument(file)
		gherkinDocuments = append(gherkinDocuments, gd)
	}

	var features []*Feature
	for _, gd := range gherkinDocuments {
		var f Feature
		f.Feature = gd.Feature
		for _, pickle := range gd.Pickles() {
			var s Scenario
			s.Pickle = pickle
			f.Scenarios = append(f.Scenarios, &s)
		}
		features = append(features, &f)
	}

	for _, f := range features {
		for _, scenario := range f.Scenarios {
			var steps []*Step
			for _, pickleStep := range scenario.Pickle.Steps {
				var step Step
				step.PickleStep = pickleStep
				steps = append(steps, &step)
			}
			scenario.Steps = steps
		}
	}
	return features
}

func filesWithExt(root, ext string) ([]string, error) {
	var filePaths []string
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ext) {
			filePaths = append(filePaths, path)
		}
		return nil
	}); err != nil {
		return make([]string, 0), err
	}
	return filePaths, nil
}