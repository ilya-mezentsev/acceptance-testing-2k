package context

import "models"

var (
	Mock = MockContext{
		Scope: map[string]interface{}{},
		ProcessingChannels: models.TestsRun{
			Success: make(chan bool),
			Error:   make(chan models.TransactionError),
		},
	}
)
