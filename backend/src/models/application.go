package models

type (
	TestsRun struct {
		Success chan bool
		Error   chan TransactionError
	}

	TransactionError struct {
		Code, Description, TransactionText string
	}
)
