package context

type MockContext struct {
	Scope map[string]interface{}
}

func (c *MockContext) ClearScope() {
	c.Scope = map[string]interface{}{}
}

func (c MockContext) SetVariable(name string, data interface{}) {
	c.Scope[name] = data
}

func (c MockContext) GetVariable(name string) interface{} {
	return c.Scope[name]
}
