package xrun

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Runner struct {
	Suite    *Suite
	Reporter *Reporter `json:"-"`
}

func (r *Runner)New() *Runner{
	r.Suite = &Suite{}
	r.Reporter = &Reporter{}
	return r
}

func (r *Runner)Run() {
	r.BuildAndRun()
	j, _ := json.MarshalIndent(r.Suite, "", "\t")
	ioutil.WriteFile("suite.json", j, os.ModePerm)
}