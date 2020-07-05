package simple

import "interfaces"

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

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	command, err := t.commandBuilder.Build(t.data.GetObject(), t.data.GetCommand())
	if err != nil {
		context.GetProcessingChannels().Error <- err
		return
	}

	_, err = command.Run(t.data.GetArguments())
	if err != nil {
		context.GetProcessingChannels().Error <- err
		return
	}

	context.GetProcessingChannels().Success <- true
}
