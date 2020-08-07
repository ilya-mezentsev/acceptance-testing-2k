package runner

type Context struct {
	Scope map[string]interface{}
}

func (c Context) SetVariable(name string, data interface{}) {
	c.Scope[name] = data
}

func (c Context) GetVariable(name string) interface{} {
	return c.Scope[name]
}
