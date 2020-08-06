package models

type (
	TestsRun struct {
		Success chan bool
		Error   chan TransactionError
	}

	ApplicationError struct {
		Code, Description string
	}

	TransactionError struct {
		Code, Description, TransactionText string
	}

	TestsReport struct {
		PassedCount, FailedCount int
		Errors                   []TransactionError
	}
)
