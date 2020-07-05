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
	command, err := t.commandBuilder.Build(t.data.GetObject(), t.data.GetCommand())
	if err != nil {
		context.GetProcessingChannels().Error <- err
		return
	}

	result, err := command.Run(t.data.GetArguments())
	if err != nil {
		context.GetProcessingChannels().Error <- err
		return
	}

	for key, value := range result {
		context.SetVariable(key, value)
	}

	context.GetProcessingChannels().Success <- true
}
