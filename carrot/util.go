package carrot

import (
	"path/filepath"
	"os"
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func printJSON(i interface{}) {
	j, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
	jf, err := json.MarshalIndent(i,"","  ")
	ioutil.WriteFile("output.json", jf, os.ModePerm)
}