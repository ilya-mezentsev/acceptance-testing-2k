package test_case

type TestCaseTransactionsIterator struct {
	currentTransactionIndex int
	testCase                string
	transactions            []string
}

func (t TestCaseTransactionsIterator) HasTransactions() bool {
	return t.currentTransactionIndex < len(t.transactions)
}

func (t *TestCaseTransactionsIterator) GetTestCaseTransaction() string {
	transactions := t.transactions[t.currentTransactionIndex]
	t.currentTransactionIndex++

	return transactions
}

func (t TestCaseTransactionsIterator) GetTestCaseText() string {
	return t.testCase
}
