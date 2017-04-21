package carrot

import (
	"os"
	"text/template"
	"log"
	"path/filepath"
	"go/build"
	"os/exec"
	"fmt"
)

const (
	mainTestFileName = "./__main/mainTest__.go"
	importFileName = "importFile__.go"
	importVarName = "Imported"
)

func WriteImportMarkers(dir string) ([]string, error) {
	const importMarkerTemplate =
		`package {{.Package}}
		var {{.VarName}} = true
		`
	var importPaths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			type ImportTemplate struct{ Package, VarName string }
			var importTemplate = ImportTemplate{info.Name(), importVarName}
			fileName := filepath.Join(path, importFileName)
			WriteTemplateToDir(fileName, importMarkerTemplate, importTemplate)

			importPath, _ := getImportPath(path)
			importPaths = append(importPaths, importPath)
		}
		return nil
	})
	return importPaths, err
}

func getImportPath(dir string) (string, error) {
	abs, _ := filepath.Abs(dir)
	pkg, err := build.ImportDir(abs, 0)
	return pkg.ImportPath, err
}

func WriteMainTestFile(importPaths []string) {
	const mainTestFileTemplate =
		`package main

		import (
			"github.com/johnmcdnl/xrun/carrot"
			{{range $n, $i := .Imports}}_i{{$n}} "{{$i}}"
			{{end}}
		)

		var (
			{{range $n, $i := .Imports}}_ = _i{{$n}}.Imported
			{{end}}

		)

		func main(){
			new(carrot.TestSuite).Run()
		}
		`
	type ImportTemplate struct{ Imports []string; VarName string }
	var importTemplate = ImportTemplate{Imports:importPaths, VarName:importVarName}
	WriteTemplateToDir(filepath.Join(mainTestFileName), mainTestFileTemplate, importTemplate)
}

func WriteTemplateToDir(outputFile string, templateText string, data interface{}) {

	t := template.Must(template.New("importTemplateMarker").Parse(templateText))

	f, err := os.Create(outputFile)
	if err != nil {
		log.Println("os.Create(filepath.Join(dir, importFileName))", err)
	}
	err = t.Execute(f, data)
	if err != nil {
		log.Println("t.Execute(f, importTemplate)", err)
	}
}

func RunMainTest(){
	cmd := exec.Command("go", "run", filepath.Join(mainTestFileName))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}