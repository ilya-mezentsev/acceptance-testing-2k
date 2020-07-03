package simple

import "interfaces"

type Transaction struct {
	repository interfaces.TransactionRepository
	data       interfaces.SimpleTransactionData
}

func New(
	repository interfaces.TransactionRepository,
	data interfaces.SimpleTransactionData,
) interfaces.Transaction {
	return Transaction{repository, data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	panic("implement me")
}
