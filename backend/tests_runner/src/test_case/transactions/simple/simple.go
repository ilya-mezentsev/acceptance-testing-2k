package simple

import (
	"interfaces"
	"models"
	"test_case/errors"
)

const (
	unableToBuildCommandError = "Unable to build simple command"
	unableToRunCommand        = "Unable to run simple command"
)

type Transaction struct {
	commandBuilder interfaces.CommandBuilder
	data           interfaces.SimpleTransactionData
}

func New(
	commandBuilder interfaces.CommandBuilder,
	data interfaces.SimpleTransactionData,
) interfaces.Transaction {
	return Transaction{commandBuilder, data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) models.TransactionError {
	command, err := t.commandBuilder.Build(t.data.GetObject(), t.data.GetCommand())
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToBuildCommandError,
			TransactionText: t.data.GetTransactionText(),
		}
	}

	_, err = command.Run(t.data.GetArguments())
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToRunCommand,
			TransactionText: t.data.GetTransactionText(),
		}
	}

	return errors.EmptyTransactionError
}
