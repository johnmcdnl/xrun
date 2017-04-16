package xrun

import (
	"path/filepath"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"go/build"
	"os/exec"
)

const importTextTemplate = `
package %s
var IMPORTED = true
`

func BuildAndRunDir(dir string) {
	defer cleanImportFiles(dir)
	cleanImportFiles(dir)
	buildDir(dir)
	runCmd()
}

func buildDir(dir string) {
	baseImportPath := baseImportPath(dir)
	importPaths := writeImportFiles(dir)
	generateMain(baseImportPath, importPaths)
}

func runCmd() {
	cmd := exec.Command("go", "run", filepath.Join(stepDefRootDir, "_xrun", "main.go"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func baseImportPath(dir string) string {
	abs, _ := filepath.Abs(dir)
	p, _ := build.ImportDir(abs, 0)
	return p.ImportPath
}

func writeImportFiles(dir string) map[string]string {
	var paths = make(map[string]string)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Name() == stepDefRootDir {
			return nil
		}
		if info.IsDir() && info.Name() != "_xrun" {
			body := fmt.Sprintf(importTextTemplate, info.Name())
			ioutil.WriteFile(fmt.Sprint(path, "/import_xrun.go"), []byte(body), os.ModePerm)
			paths[info.Name()] = filepath.ToSlash(path)
		}
		return nil
	})
	return paths
}

func cleanImportFiles(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "_xrun") {
			os.RemoveAll(path)
		}
		return nil
	})
}

func generateMain(baseImportPath string, imports map[string]string) error {
	//TODO - this may need some cleanup :-P
	var body string
	body = fmt.Sprint(body, "package main \n")
	body = fmt.Sprint(body, "import ( \n")
	for _, path := range imports {
		path = strings.Join(strings.SplitAfter(path, "/")[1:], "")
		body = fmt.Sprint(body, fmt.Sprintf(`"%s/%s"`, baseImportPath, path), "\n")
	}
	body = fmt.Sprint(body, `"github.com/johnmcdnl/xrun/xrun"`)
	body = fmt.Sprint(body, ")\n")

	body = fmt.Sprint(body, `var (`, "\n")
	for name, _ := range imports {
		body = fmt.Sprint(body, fmt.Sprintf(`_ = %s.IMPORTED`, name), "\n")
	}
	body = fmt.Sprint(body, `)`, "\n")

	body = fmt.Sprint(body, `func main() {`, "\n")
	body = fmt.Sprint(body, `new(xrun.Runner).New().Run()`, "\n")
	body = fmt.Sprint(body, `}`, "\n")

	err := os.MkdirAll(filepath.Join(stepDefRootDir, "_xrun"), os.ModePerm)
	if err != nil {
		//	return err
	}
	err = ioutil.WriteFile(filepath.Join(stepDefRootDir, "_xrun", "main.go"), []byte(body), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
