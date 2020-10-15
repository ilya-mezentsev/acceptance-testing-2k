package assertion

import (
	"test_case/errors"
	"test_case/transactions/plugins/value_path"
	mockAssertion "test_runner_meta/mock/transaction/assertion"
	mockContext "test_runner_meta/mock/transaction/context"
	"test_utils"
	"testing"
)

var context = mockContext.Mock

func TestTransaction_ExecuteSuccess(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": map[string]interface{}{"score": "10"},
	})
	transaction := New(&mockAssertion.MockDataScore10)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
}

func TestTransaction_ExecuteSuccessAssertionFailed(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": map[string]interface{}{"score": "11"},
	})
	transaction := New(&mockAssertion.MockDataScore10)

	err := transaction.Execute(context)

	test_utils.AssertEqual(assertionFailed.Error(), err.Code, t)
	test_utils.AssertEqual("Expected: 10, but got: 11", err.Description, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTestCaseText(), err.TestCaseText, t)
}

func TestTransaction_ExecuteSuccessTemplateReplacement(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": map[string]interface{}{"score": "11"},
	})
	context.SetVariable("foo", map[string]interface{}{
		"bar": "11",
	})
	transaction := New(&mockAssertion.MockDataTemplateNewValue)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
}

func TestTransaction_ExecuteFailedTemplateReplacement(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": map[string]interface{}{"score": "11"},
	})
	transaction := New(&mockAssertion.MockDataTemplateNewValue)

	err := transaction.Execute(context)

	test_utils.AssertEqual(value_path.CannotAccessValueByPath.Error(), err.Code, t)
	test_utils.AssertEqual("Unable to process new value: ${foo.bar}", err.Description, t)
	test_utils.AssertEqual(mockAssertion.MockDataTemplateNewValue.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(mockAssertion.MockDataTemplateNewValue.GetTestCaseText(), err.TestCaseText, t)
}

func TestTransaction_ExecuteSuccessArrayValue(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": []interface{}{"1", "2", "3"},
	})
	transaction := New(&mockAssertion.MockDataArray)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
}

func TestTransaction_ExecuteSuccessArrayWithMap(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": []interface{}{map[string]interface{}{
			"x": "1", "y": "2",
		}},
	})
	transaction := New(&mockAssertion.MockDataArrayWithMap)

	err := transaction.Execute(context)

	test_utils.AssertEqual(errors.EmptyTransactionError, err, t)
}

func TestTransaction_ExecuteCannotAccessValue(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": "[1, 2, 3]",
	})
	transaction := New(&mockAssertion.MockDataScore10)

	err := transaction.Execute(context)
	test_utils.AssertEqual(value_path.CannotAccessValueByPath.Error(), err.Code, t)
	test_utils.AssertEqual(
		"Unable to get value by path: "+mockAssertion.MockDataScore10.GetDataPath(),
		err.Description,
		t,
	)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTestCaseText(), err.TestCaseText, t)
}

func TestTransaction_ExecuteAssertionFailedByTypes(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": []interface{}{map[string]interface{}{
			"x": "1", "y": []interface{}{"0", "1"},
		}},
	})
	transaction := New(&mockAssertion.MockDataArrayWithMap)

	err := transaction.Execute(context)
	test_utils.AssertEqual(assertionFailed.Error(), err.Code, t)
	test_utils.AssertEqual("Expected: 2, but got: [0 1]", err.Description, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTestCaseText(), err.TestCaseText, t)
}

func TestTransaction_ExecuteVariableIsNotDefined(t *testing.T) {
	transaction := New(&mockAssertion.MockDataScore10)

	err := transaction.Execute(context)
	test_utils.AssertEqual(variableIsNotDefined.Error(), err.Code, t)
	test_utils.AssertEqual("Unable to find variable: response", err.Description, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
	test_utils.AssertEqual(mockAssertion.MockDataScore10.GetTestCaseText(), err.TestCaseText, t)
}
