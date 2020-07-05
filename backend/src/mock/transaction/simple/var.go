package simple

import (
	"errors"
	"models"
	"test_case/parsers/transaction/data"
)

var (
	BuildCommandError = errors.New("build-error")
	RunCommandError   = errors.New("run-error")
	MockData          data.SimpleTransactionData
	Context           = MockContext{
		ProcessingChannels: models.TestsRun{
			Success: make(chan bool),
			Result:  make(chan models.TestResult),
			Error:   make(chan error),
		},
	}
)
