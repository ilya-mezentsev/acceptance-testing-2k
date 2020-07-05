package assertion

import (
	"interfaces"
)

type Transaction struct {
	commandBuilder interfaces.CommandBuilder
	data           interfaces.AssertionTransactionData
}

func New(
	commandBuilder interfaces.CommandBuilder,
	data interfaces.AssertionTransactionData,
) interfaces.Transaction {
	return Transaction{commandBuilder, data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	panic("implement me")
}
