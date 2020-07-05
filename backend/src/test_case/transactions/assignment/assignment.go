package assignment

import "interfaces"

type Transaction struct {
	commandBuilder interfaces.CommandBuilder
	data           interfaces.AssignmentTransactionData
}

func New(
	commandBuilder interfaces.CommandBuilder,
	data interfaces.AssignmentTransactionData,
) interfaces.Transaction {
	return Transaction{commandBuilder, data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	panic("implement me")
}
