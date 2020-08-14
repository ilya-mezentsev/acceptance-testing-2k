package simple

import (
	"test_case/errors"
	mockConst "test_runner_meta/mock/transaction/constant"
	mockContext "test_runner_meta/mock/transaction/context"
	"test_runner_meta/mock/transaction/simple"
	"test_utils"
	"testing"
)

var context = mockContext.Mock

func TestTransaction_ExecuteSuccess(t *testing.T) {
	transaction := New(
		simple.MockCommandBuilder{},
		&simple.MockData,
	)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
}

func TestTransaction_ExecuteBuildCommandError(t *testing.T) {
	transaction := New(
		simple.MockCommandBuilderWithError{},
		&simple.MockData,
	)

	err := transaction.Execute(context)

	test_utils.AssertEqual(mockConst.BuildCommandError.Error(), err.Code, t)
	test_utils.AssertEqual(unableToBuildCommandError, err.Description, t)
	test_utils.AssertEqual(simple.MockData.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(simple.MockData.GetTestCaseText(), err.TestCaseText, t)
}

func TestTransaction_ExecuteCommandRunError(t *testing.T) {
	transaction := New(
		simple.MockErroredCommandBuilder{},
		&simple.MockData,
	)

	err := transaction.Execute(context)

	test_utils.AssertEqual(mockConst.RunCommandError.Error(), err.Code, t)
	test_utils.AssertEqual(unableToRunCommand, err.Description, t)
	test_utils.AssertEqual(simple.MockData.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(simple.MockData.GetTestCaseText(), err.TestCaseText, t)
}
