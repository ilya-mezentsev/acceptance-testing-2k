package interfaces

type (
	TestCasesIterator interface {
		Init(testCases string) error
		Done() bool
		NextTransactions() []string
	}
)
