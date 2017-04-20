package carrot

type Result struct {
	IsExecuted bool
	IsPassed   bool
	Errors     []error
}