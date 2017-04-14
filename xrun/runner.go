package xrun

type Runner struct {
	Builder  Builder
	Suite    Suite
	Reporter Reporter
}

func (r *Runner)Run() {
	r.Builder.Run()
	r.Suite.Run()
	r.Reporter.Run()
}