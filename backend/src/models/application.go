package models

type (
	TestsRun struct {
		Success chan bool
		Result  chan TestResult
		Error   chan error
	}

	TestResult struct {
	}
)
