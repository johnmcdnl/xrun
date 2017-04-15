package xrun

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"github.com/cucumber/gherkin-go"
	"io/ioutil"
	"os/exec"
	"text/template"
	"encoding/json"
	"sync"
)

type Builder struct {

}

const featuresDir = "./internal/features"
const stepDefDir = "./internal"

func (r *Runner)Build() {
	fmt.Println("START (r *Runner)Build()")
	os.MkdirAll(featuresDir, os.ModePerm)
	os.MkdirAll(stepDefDir, os.ModePerm)
	fmt.Println("DIR's CREATED")
	b := Builder{}
	s := NewSuite()
	fmt.Println("ANOTHER LINE HITS THE DUST")
	s.Features = b.gherkinToFeatures()
	s.StepDefs = b.GetStepDefs()
	r.Suite = s

	fmt.Println("FINISHED (r *Runner)Build()")

	r.Suite.Run()
	r.Reporter.Run()

	j, _ := json.MarshalIndent(r.Suite, "", "\t")
	ioutil.WriteFile("suite.json", j, os.ModePerm)
}

var buildOnce sync.Once

func (b *Builder)GetStepDefs() []*StepDef {
	//goFiles, _ := filesWithExt(stepDefDir, ".go")
	//buildAndRunDir(stepDefDir, goFiles, []string{}, "")

	return []*StepDef{}
}

func BuildAndRun(){
	goFiles, _ := filesWithExt(stepDefDir, ".go")
	buildAndRunDir(stepDefDir, goFiles, []string{}, "")
}

func buildAndRunDir(dir string, goFiles []string, filters []string, goBuildTags string) error {
	buildCleanup(dir)
	defer buildCleanup(dir)

	info := buildInfo{
		Imports:      []string{},
		FeaturesPath: fmt.Sprintf("%q", dir),
		Filters:      filters,
	}


	// write special constants to packages so they can be imported
	for _, file := range goFiles {
		ifile := filepath.Join(filepath.Dir(file), importMarkerFile)
		if _, err := os.Stat(ifile); err != nil {
			pkgName := filepath.Base(filepath.Dir(file))
			if pkgName == "_test" {
				continue
			}
			fullPkg := assembleImportPath(file)

			if fullPkg == "" {
				return fmt.Errorf("could not determine package path for %s", file)
			}

			info.Imports = append(info.Imports, fullPkg)

			src := fmt.Sprintf("package %s\nvar IMPORT_MARKER = true\n", pkgName)
			err = ioutil.WriteFile(ifile, []byte(src), 0664)
			if err != nil {
				return err
			}
		}
	}

	// write main test stub
	os.MkdirAll(filepath.Join(dir, "_test"), 0777)
	f, err := os.Create(filepath.Join(dir, "_test", testFile))
	if err != nil {
		return err
	}
	tplMain.Execute(f, info)
	f.Close()

	// now run the command
	tfile := "./" + filepath.ToSlash(dir) + "/_test/" + testFile
	var cmd *exec.Cmd
	if len(goBuildTags) > 0 {
		cmd = exec.Command("go", "run", "-tags", goBuildTags, tfile)
	} else {
		cmd = exec.Command("go", "run", tfile)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		buildCleanup(dir)
		os.Exit(1)
	}

	return nil
}

// ToSlash is being used to coerce the different
// os PathSeparators into the forward slash
// as the forward slash is required by Go's import statement
func assembleImportPath(file string) string {
	a, _ := filepath.Abs(filepath.Dir(file))
	absPath, fullPkg := filepath.ToSlash(a), ""
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		a, _ = filepath.Abs(p)
		p = filepath.ToSlash(a)
		if strings.HasPrefix(absPath, p) {
			prefixPath := filepath.ToSlash(filepath.Join(p, "src"))
			rpath, _ := filepath.Rel(prefixPath, absPath)
			fullPkg = filepath.ToSlash(rpath)
			break
		}
	}
	return fullPkg
}

type buildInfo struct {
	Imports      []string
	FeaturesPath string
	Filters      []string
}

const (
	importMarkerFile = "importmarker__.go"
	testFile = "xruncoretest__.go"
)

func buildCleanup(dir string) {

	importMarkerFiles, _ := filesWithExt(stepDefDir, importMarkerFile)
	testFile, _ := filesWithExt(stepDefDir, testFile)
	for _, file := range importMarkerFiles {
		os.Remove(file)
	}

	for _, file := range testFile {
		os.Remove(file)
	}

	p := filepath.Join(dir, "_test")
	if _, err := os.Stat(p); err == nil {
		os.RemoveAll(p)
	}
}

var tplMain = template.Must(template.New("main").Parse(`
package main

import (
	"github.com/johnmcdnl/xrun/xrun"
	{{range $n, $i := .Imports}}_i{{$n}} "{{$i}}"
	{{end}}
)

var (
	{{range $n, $i := .Imports}}_ci{{$n}} = _i{{$n}}.IMPORT_MARKER
	{{end}}
)

func main() {
	new(xrun.Runner).New().Run()
}
`))

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