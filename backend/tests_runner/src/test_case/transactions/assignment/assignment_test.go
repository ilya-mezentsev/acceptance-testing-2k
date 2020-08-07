package assignment

import (
	"mock/transaction/assignment"
	"mock/transaction/constant"
	mockContext "mock/transaction/context"
	"test_case/errors"
	"test_utils"
	"testing"
)

var context = mockContext.Mock

func TestTransaction_ExecuteNilResultCommand(t *testing.T) {
	transaction := New(
		assignment.MockNilResultCommandBuilder{},
		&assignment.MockData,
	)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
	test_utils.AssertEqual(
		0,
		len(
			context.GetVariable(
				assignment.MockData.GetVariableName()).(map[string]interface{})),
		t,
	)
}

func TestTransaction_ExecuteNotNilResultCommand(t *testing.T) {
	transaction := New(
		assignment.MockNotNilResultCommandBuilder{},
		&assignment.MockData,
	)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
	for key, value := range assignment.MockCommandResult {
		test_utils.AssertEqual(
			value,
			context.GetVariable(
				assignment.MockData.GetVariableName()).(map[string]interface{})[key],
			t,
		)
	}
}

func TestTransaction_ExecuteBuildCommandError(t *testing.T) {
	transaction := New(
		assignment.MockCommandBuilderWithError{},
		&assignment.MockData,
	)

	err := transaction.Execute(context)
	test_utils.AssertEqual(constant.BuildCommandError.Error(), err.Code, t)
	test_utils.AssertEqual(unableToBuildCommand, err.Description, t)
	test_utils.AssertEqual(assignment.MockData.GetTransactionText(), err.TransactionText, t)
}

func TestTransaction_ExecuteCommandRunError(t *testing.T) {
	transaction := New(
		assignment.MockErroredCommandBuilder{},
		&assignment.MockData,
	)

	err := transaction.Execute(context)
	test_utils.AssertEqual(constant.RunCommandError.Error(), err.Code, t)
	test_utils.AssertEqual(unableToRunCommand, err.Description, t)
	test_utils.AssertEqual(assignment.MockData.GetTransactionText(), err.TransactionText, t)
}
