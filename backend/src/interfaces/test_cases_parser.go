package interfaces

type (
	TestCasesParser interface {
		Parse(testCases string) ([]TestCaseTransactionsIterator, error)
	}

	TestCaseTransactionsIterator interface {
		HasTransactions() bool
		GetTestCaseTransaction() string
	}
)
