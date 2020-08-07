package errors

import (
	"errors"
	"models"
)

var (
	UnknownTransactionType   = errors.New("unknown-transaction-type")
	NoTransactionsInTestCase = errors.New("no-transactions-in-test-case")
)

var EmptyTransactionError models.TransactionError
