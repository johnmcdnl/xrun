package carrot

type TestContext struct {
	Id       string
	data     map[string]interface{}
	TestingT interface{}
}

func NewContext(id string) *TestContext {
	var tc TestContext
	tc.Id = id
	tc.data = make(map[string]interface{})
	return &tc
}

func (tc *TestContext)Add(key string, value interface{}) {
	tc.data[key] = value
}

func (tc *TestContext)Get(key string) interface{} {
	return tc.data[key]
}