package assertion

import (
	"fmt"
	"interfaces"
	"models"
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
		context.GetProcessingChannels().Error <- models.TransactionError{
			Code:            variableIsNotDefined.Error(),
			Description:     t.variableIsNotDefinedDescription(),
			TransactionText: t.data.GetTransactionText(),
		}
		return
	}

	currentValue, err := getValueByPath(
		context.GetVariable(t.data.GetVariableName()),
		t.data.GetDataPath(),
	)
	if err != nil {
		context.GetProcessingChannels().Error <- models.TransactionError{
			Code:            err.Error(),
			Description:     t.unableToGetValueByPathDescription(),
			TransactionText: t.data.GetTransactionText(),
		}
		return
	}

	if t.equals(currentValue, t.data.GetNewValue()) {
		context.GetProcessingChannels().Success <- true
	} else {
		context.GetProcessingChannels().Error <- models.TransactionError{
			Code:            assertionFailed.Error(),
			Description:     t.assertionFailedDescription(currentValue),
			TransactionText: t.data.GetTransactionText(),
		}
	}
}

func (t Transaction) variableExists(context interfaces.TestCaseContext) bool {
	return context.GetVariable(t.data.GetVariableName()) != nil
}

func (t Transaction) variableIsNotDefinedDescription() string {
	return fmt.Sprintf("Unable to find variable: %s", t.data.GetVariableName())
}

func (t Transaction) unableToGetValueByPathDescription() string {
	return fmt.Sprintf("Unable to get value by path: %s", t.data.GetDataPath())
}

func (t Transaction) assertionFailedDescription(current interface{}) string {
	return fmt.Sprintf("Expected: %v, but got: %v", t.data.GetNewValue(), current)
}

func (t Transaction) equals(current interface{}, expected string) bool {
	strCurrent, ok := current.(string)

	return ok && strCurrent == expected
}
