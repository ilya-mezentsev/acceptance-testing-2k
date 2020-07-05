package context

import (
	"models"
)

type MockContext struct {
	Scope              map[string]interface{}
	ProcessingChannels models.TestsRun
}

func (c MockContext) SetVariable(name string, data interface{}) {
	c.Scope[name] = data
}

func (c MockContext) GetVariable(name string) interface{} {
	return c.Scope[name]
}

func (c MockContext) GetProcessingChannels() models.TestsRun {
	return c.ProcessingChannels
}
