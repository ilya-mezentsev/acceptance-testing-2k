package assignment

import (
	"test_case/errors"
	"test_case/transactions/plugins/arguments_processor"
	"test_runner_meta/interfaces"
	"test_runner_meta/models"
)

const (
	unableToBuildCommand = "Unable to build assignment command"
	unableToRunCommand   = "Unable to run assignment command"
)

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

func (t Transaction) Execute(context interfaces.TestCaseContext) models.TransactionError {
	command, err := t.commandBuilder.Build(t.data.GetObject(), t.data.GetCommand())
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToBuildCommand,
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	arguments, err := arguments_processor.ReplaceTemplatesWithVariables(
		context,
		t.data.GetArguments(),
	)
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToRunCommand,
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	result, err := command.Run(arguments)
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     unableToRunCommand,
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	context.SetVariable(t.data.GetVariableName(), result)
	return errors.EmptyTransactionError
}
