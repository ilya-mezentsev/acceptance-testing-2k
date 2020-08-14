package assertion

import (
	"fmt"
	"test_case/errors"
	"test_runner_meta/interfaces"
	"test_runner_meta/models"
)

type Transaction struct {
	data interfaces.AssertionTransactionData
}

func New(
	data interfaces.AssertionTransactionData,
) interfaces.Transaction {
	return Transaction{data}
}

func (t Transaction) Execute(context interfaces.TestCaseContext) models.TransactionError {
	if !t.variableExists(context) {
		return models.TransactionError{
			Code:            variableIsNotDefined.Error(),
			Description:     t.variableIsNotDefinedDescription(),
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	currentValue, err := getValueByPath(
		context.GetVariable(t.data.GetVariableName()),
		t.data.GetDataPath(),
	)
	if err != nil {
		return models.TransactionError{
			Code:            err.Error(),
			Description:     t.unableToGetValueByPathDescription(),
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
		}
	}

	if t.equals(currentValue, t.data.GetNewValue()) {
		return errors.EmptyTransactionError
	} else {
		return models.TransactionError{
			Code:            assertionFailed.Error(),
			Description:     t.assertionFailedDescription(currentValue),
			TransactionText: t.data.GetTransactionText(),
			TestCaseText:    t.data.GetTestCaseText(),
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
