package simple

import (
	"test_case/errors"
	"test_runner_meta/interfaces"
	"test_runner_meta/models"
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

func (t Transaction) Execute(interfaces.TestCaseContext) models.TransactionError {
	command, err := t.commandBuilder.Build(t.data.GetObject(), t.data.GetCommand())
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToBuildCommandError,
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	_, err = command.Run(t.data.GetArguments())
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToRunCommand,
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	return errors.EmptyTransactionError
}
