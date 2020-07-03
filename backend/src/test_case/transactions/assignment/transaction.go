package assignment

import "interfaces"

type Transaction struct {
	repository interfaces.TransactionRepository
	data       interfaces.AssignmentTransactionData
}

func New(
	repository interfaces.TransactionRepository,
	data interfaces.AssignmentTransactionData,
) interfaces.Transaction {
	return Transaction{repository, data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	panic("implement me")
}
