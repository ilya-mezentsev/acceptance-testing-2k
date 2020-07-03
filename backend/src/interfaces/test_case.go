package interfaces

import "models"

type (
	TestCase interface {
		Run(processing models.TestsRun)
	}

	Transaction interface {
		Execute(context TestCaseContext)
	}

	TestCaseContext interface {
		GetVariable(name string) map[string]interface{}
		GetProcessingChannels() models.TestsRun
	}
)
