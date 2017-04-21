package carrot

import "sync"

var (
	testfeatureSync sync.WaitGroup
	testCaseSync sync.WaitGroup
	testStepSync sync.WaitGroup
)
