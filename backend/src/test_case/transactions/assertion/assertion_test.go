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
			utils.AssertErrorsEqual(assertionFailed, err, t)
			return
		}
	}
}

func TestTransaction_ExecuteSuccessArrayValue(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("response", map[string]interface{}{
		"data": "[1, 2, 3]",
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
			utils.AssertErrorsEqual(cannotAccessValueByPath, err, t)
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
			utils.AssertErrorsEqual(variableIsNotDefined, err, t)
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

func TestTransaction_IsMapTrue(t *testing.T) {
	utils.AssertTrue(isMap(map[string]interface{}{}), t)
}

func TestTransaction_IsMapFalse(t *testing.T) {
	utils.AssertFalse(isMap(""), t)
}
