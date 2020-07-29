package simple

import (
	mockConst "mock/transaction/constant"
	mockContext "mock/transaction/context"
	"mock/transaction/simple"
	"test_utils"
	"testing"
)

var context = mockContext.Mock

func TestTransaction_ExecuteSuccess(t *testing.T) {
	transaction := New(
		simple.MockCommandBuilder{},
		&simple.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case err := <-context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
		case res := <-context.GetProcessingChannels().Success:
			test_utils.AssertTrue(res, t)
			return
		}
	}
}

func TestTransaction_ExecuteBuildCommandError(t *testing.T) {
	transaction := New(
		simple.MockCommandBuilderWithError{},
		&simple.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case err := <-context.GetProcessingChannels().Error:
			test_utils.AssertEqual(mockConst.BuildCommandError.Error(), err.Code, t)
			test_utils.AssertEqual(unableToBuildCommandError, err.Description, t)
			test_utils.AssertEqual(simple.MockData.GetTransactionText(), err.TransactionText, t)
			return
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
		}
	}
}

func TestTransaction_ExecuteCommandRunError(t *testing.T) {
	transaction := New(
		simple.MockErroredCommandBuilder{},
		&simple.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case err := <-context.GetProcessingChannels().Error:
			test_utils.AssertEqual(mockConst.RunCommandError.Error(), err.Code, t)
			test_utils.AssertEqual(unableToRunCommand, err.Description, t)
			test_utils.AssertEqual(simple.MockData.GetTransactionText(), err.TransactionText, t)
			return
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
		}
	}
}
