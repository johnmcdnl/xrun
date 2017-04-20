package xrun

type TestContext struct {
	ID       string
	Data     map[string]interface{}
	TestingT *TestingT
}

func testContextWithId(id string) *TestContext {
	var tc TestContext
	tc.ID = id
	tc.Data = make(map[string]interface{})
	tc.TestingT = &TestingT{}
	return &tc
}

func (tc TestContext)T() *TestingT {
	return tc.TestingT
}
