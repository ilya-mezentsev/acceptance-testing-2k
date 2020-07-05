package simple

import (
	"models"
)

type MockContext struct {
	ProcessingChannels models.TestsRun
}

func (c MockContext) SetVariable(string, map[string]interface{}) {
	panic("implement me")
}

func (c MockContext) GetVariable(string) map[string]interface{} {
	panic("implement me")
}

func (c MockContext) GetProcessingChannels() models.TestsRun {
	return c.ProcessingChannels
}
