package assignment

import (
	"interfaces"
	"models"
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

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	command, err := t.commandBuilder.Build(t.data.GetObject(), t.data.GetCommand())
	if err != nil {
		context.GetProcessingChannels().Error <- models.TransactionError{
			Code:            err.Error(),
			Description:     unableToBuildCommand,
			TransactionText: t.data.GetTransactionText(),
		}
		return
	}

	result, err := command.Run(t.data.GetArguments())
	if err != nil {
		context.GetProcessingChannels().Error <- models.TransactionError{
			Code:            err.Error(),
			Description:     unableToRunCommand,
			TransactionText: t.data.GetTransactionText(),
		}
		return
	}

	context.SetVariable(t.data.GetVariableName(), result)
	context.GetProcessingChannels().Success <- true
}
