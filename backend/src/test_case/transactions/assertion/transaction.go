package assertion

import (
	"interfaces"
)

type Transaction struct {
	repository interfaces.TransactionRepository
	data       interfaces.AssertionTransactionData
}

func New(
	repository interfaces.TransactionRepository,
	data interfaces.AssertionTransactionData,
) interfaces.Transaction {
	return Transaction{repository, data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	panic("implement me")
}
