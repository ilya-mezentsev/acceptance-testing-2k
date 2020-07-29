package errors

import "errors"

var (
	UnknownTransactionType   = errors.New("unknown-transaction-type")
	NoTransactionsInTestCase = errors.New("no-transactions-in-test-case")
)
