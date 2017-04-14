package xrun

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Runner struct {
	Builder  *Builder
	Suite    *Suite
	Reporter *Reporter
}

func (r *Runner)New() *Runner{
	r.Builder = &Builder{}
	r.Suite = &Suite{}
	r.Reporter = &Reporter{}
	return r
}

func (r *Runner)Run() {
	r.Build()
	r.Suite.Run()
	r.Reporter.Run()

	j, _ := json.MarshalIndent(r.Suite, "", "\t")
	ioutil.WriteFile("suite.json", j, os.ModePerm)
}