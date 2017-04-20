package carrot

type TestContext struct {
	Id   string
	data map[string]interface{}
}

func (tc *TestContext)Add(key string, value interface{}) {
	tc.data[key] = value
}

func (tc *TestContext)Get(key string) interface{} {
	return tc.data[key]
}