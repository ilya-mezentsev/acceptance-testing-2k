package assertion

import (
	mockAssertion "mock/transaction/assertion"
	mockContext "mock/transaction/context"
	"testing"
	"utils"
)

var context = mockContext.Mock

func TestTransaction_ExecuteSuccess(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": map[string]interface{}{"score": "10"},
	})
	transaction := New(&mockAssertion.MockDataScore10)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			return
		case err := <-context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
			return
		}
	}
}

func TestTransaction_ExecuteSuccessAssertionFailed(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": map[string]interface{}{"score": "11"},
	})
	transaction := New(&mockAssertion.MockDataScore10)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-context.GetProcessingChannels().Error:
			utils.AssertEqual(assertionFailed.Error(), err.Code, t)
			utils.AssertEqual("Expected: 10, but got: 11", err.Description, t)
			utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
			return
		}
	}
}

func TestTransaction_ExecuteSuccessArrayValue(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": []interface{}{"1", "2", "3"},
	})
	transaction := New(&mockAssertion.MockDataArray)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			return
		case err := <-context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
			return
		}
	}
}

func TestTransaction_ExecuteSuccessArrayWithMap(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": []interface{}{map[string]interface{}{
			"x": "1", "y": "2",
		}},
	})
	transaction := New(&mockAssertion.MockDataArrayWithMap)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			return
		case err := <-context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
			return
		}
	}
}

func TestTransaction_ExecuteCannotAccessValue(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": "[1, 2, 3]",
	})
	transaction := New(&mockAssertion.MockDataScore10)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-context.GetProcessingChannels().Error:
			utils.AssertEqual(cannotAccessValueByPath.Error(), err.Code, t)
			utils.AssertEqual("Unable to get value by path: "+mockAssertion.MockDataScore10.GetDataPath(), err.Description, t)
			utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
			return
		}
	}
}

func TestTransaction_ExecuteAssertionFailedByTypes(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": []interface{}{map[string]interface{}{
			"x": "1", "y": []interface{}{"0", "1"},
		}},
	})
	transaction := New(&mockAssertion.MockDataArrayWithMap)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-context.GetProcessingChannels().Error:
			utils.AssertEqual(assertionFailed.Error(), err.Code, t)
			utils.AssertEqual("Expected: 2, but got: [0 1]", err.Description, t)
			utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
			return
		}
	}
}

func TestTransaction_ExecuteVariableIsNotDefined(t *testing.T) {
	transaction := New(&mockAssertion.MockDataScore10)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-context.GetProcessingChannels().Error:
			utils.AssertEqual(variableIsNotDefined.Error(), err.Code, t)
			utils.AssertEqual("Unable to find variable: response", err.Description, t)
			utils.AssertEqual(mockAssertion.MockDataScore10.GetTransactionText(), err.TransactionText, t)
			return
		}
	}
}

func TestTransaction_GetValueByPathSingleKey(t *testing.T) {
	value, err := getValueByPath(map[string]interface{}{
		"x": 10,
	}, "x")

	utils.AssertNil(err, t)
	utils.AssertEqual(10, value.(int), t)
}

func TestTransaction_GetValueByPathDotSeparated(t *testing.T) {
	value, err := getValueByPath(map[string]interface{}{
		"x": map[string]interface{}{
			"y": 10,
		},
	}, "x.y")

	utils.AssertNil(err, t)
	utils.AssertEqual(10, value.(int), t)
}

func TestTransaction_GetValueByPathArray(t *testing.T) {
	value, err := getValueByPath(map[string]interface{}{
		"x": []interface{}{1},
	}, "x.0")

	utils.AssertNil(err, t)
	utils.AssertEqual(1, value.(int), t)
}

func TestTransaction_GetValueByPathArrayWithMap(t *testing.T) {
	value, err := getValueByPath(map[string]interface{}{
		"x": []interface{}{map[string]interface{}{
			"y": 1,
		}},
	}, "x.0.y")

	utils.AssertNil(err, t)
	utils.AssertEqual(1, value.(int), t)
}

func TestTransaction_GetValueByPathArrayIndexOutOfBounds(t *testing.T) {
	_, err := getValueByPath(map[string]interface{}{
		"x": []interface{}{1},
	}, "x.1")

	utils.AssertErrorsEqual(indexOutOfBounds, err, t)
}

func TestTransaction_GetValueByPathArrayInvalidIndex(t *testing.T) {
	_, err := getValueByPath(map[string]interface{}{
		"x": []interface{}{1},
	}, "x.a")

	utils.AssertErrorsEqual(invalidNumberForIndex, err, t)
}

func TestTransaction_GetValueByPathInvalidPath(t *testing.T) {
	_, err := getValueByPath(map[string]interface{}{
		"x": 1,
	}, "")

	utils.AssertErrorsEqual(invalidPath, err, t)
}

func TestTransaction_GetValueByPathInvalidValue(t *testing.T) {
	_, err := getValueByPath(10, "x")

	utils.AssertErrorsEqual(cannotAccessValueByPath, err, t)
}
