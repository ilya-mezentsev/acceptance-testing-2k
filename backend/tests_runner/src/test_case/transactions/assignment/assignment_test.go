package assignment

import (
	"mock/transaction/assignment"
	"mock/transaction/constant"
	mockContext "mock/transaction/context"
	"test_utils"
	"testing"
)

var context = mockContext.Mock

func TestTransaction_ExecuteNilResultCommand(t *testing.T) {
	transaction := New(
		assignment.MockNilResultCommandBuilder{},
		&assignment.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			test_utils.AssertEqual(
				0,
				len(
					context.GetVariable(
						assignment.MockData.GetVariableName()).(map[string]interface{})),
				t,
			)
			return
		case err := <-context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
			return
		}
	}
}

func TestTransaction_ExecuteNotNilResultCommand(t *testing.T) {
	transaction := New(
		assignment.MockNotNilResultCommandBuilder{},
		&assignment.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case <-context.GetProcessingChannels().Success:
			for key, value := range assignment.MockCommandResult {
				test_utils.AssertEqual(
					value,
					context.GetVariable(
						assignment.MockData.GetVariableName()).(map[string]interface{})[key],
					t,
				)
			}
			return
		case err := <-context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
			return
		}
	}
}

func TestTransaction_ExecuteBuildCommandError(t *testing.T) {
	transaction := New(
		assignment.MockCommandBuilderWithError{},
		&assignment.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case err := <-context.GetProcessingChannels().Error:
			test_utils.AssertEqual(constant.BuildCommandError.Error(), err.Code, t)
			test_utils.AssertEqual(unableToBuildCommand, err.Description, t)
			test_utils.AssertEqual(assignment.MockData.GetTransactionText(), err.TransactionText, t)
			return
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
		}
	}
}

func TestTransaction_ExecuteCommandRunError(t *testing.T) {
	transaction := New(
		assignment.MockErroredCommandBuilder{},
		&assignment.MockData,
	)

	go transaction.Execute(context)

	for {
		select {
		case err := <-context.GetProcessingChannels().Error:
			test_utils.AssertEqual(constant.RunCommandError.Error(), err.Code, t)
			test_utils.AssertEqual(unableToRunCommand, err.Description, t)
			test_utils.AssertEqual(assignment.MockData.GetTransactionText(), err.TransactionText, t)
			return
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
		}
	}
}
