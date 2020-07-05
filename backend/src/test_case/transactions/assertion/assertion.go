package assertion

import (
	"interfaces"
)

type Transaction struct {
	data interfaces.AssertionTransactionData
}

func New(
	data interfaces.AssertionTransactionData,
) interfaces.Transaction {
	return Transaction{data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) {
	if !t.variableExists(context) {
		context.GetProcessingChannels().Error <- variableIsNotDefined
		return
	}

	currentValue, err := getValueByPath(
		context.GetVariable(t.data.GetVariableName()),
		t.data.GetDataPath(),
	)
	if err != nil {
		context.GetProcessingChannels().Error <- err
		return
	}

	if t.equals(currentValue, t.data.GetNewValue()) {
		context.GetProcessingChannels().Success <- true
	} else {
		context.GetProcessingChannels().Error <- assertionFailed
	}
}

func (t Transaction) variableExists(context interfaces.TestCaseContext) bool {
	return context.GetVariable(t.data.GetVariableName()) != nil
}

func (t Transaction) equals(current interface{}, expected string) bool {
	strCurrent, ok := current.(string)

	return ok && strCurrent == expected
}
