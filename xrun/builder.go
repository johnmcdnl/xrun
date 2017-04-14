package xrun

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"github.com/cucumber/gherkin-go"
)

type Builder struct {

}

const featuresDir = "./internal/features"
const stepDefDir = "./internal"

func (r *Runner)Build() {
	os.MkdirAll(featuresDir, os.ModePerm)
	os.MkdirAll(stepDefDir, os.ModePerm)

	b := Builder{}
	s := NewSuite()

	s.Features = b.gherkinToFeatures()
	s.StepDefs = b.GetStepDefs()
	r.Suite = s

	fmt.Println("(b *Builder)Run()")
}

func (b *Builder)gherkinToFeatures() []*Feature {
	featureFiles, _ := filesWithExt(featuresDir, ".feature")

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

func (b *Builder)GetStepDefs() []*StepDef {

	return []*StepDef{}
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