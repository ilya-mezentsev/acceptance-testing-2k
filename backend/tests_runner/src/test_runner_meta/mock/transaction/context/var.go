package context

var (
	Mock = MockContext{
		Scope: map[string]interface{}{},
	}
)
