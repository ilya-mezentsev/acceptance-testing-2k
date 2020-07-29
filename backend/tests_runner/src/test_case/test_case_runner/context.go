package test_case_runner

import "models"

type Context struct {
	Scope              map[string]interface{}
	ProcessingChannels models.TestsRun
}

func (c Context) SetVariable(name string, data interface{}) {
	c.Scope[name] = data
}

func (c Context) GetVariable(name string) interface{} {
	return c.Scope[name]
}

func (c Context) GetProcessingChannels() models.TestsRun {
	return c.ProcessingChannels
}
