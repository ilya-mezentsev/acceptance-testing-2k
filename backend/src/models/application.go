package models

type (
	TestsRun struct {
		Success chan bool
		Error   chan error
	}
)
