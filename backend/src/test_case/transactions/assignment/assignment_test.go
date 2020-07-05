package assignment

import (
	"mock/transaction/assignment"
	"mock/transaction/constant"
	mockContext "mock/transaction/context"
	"testing"
	"utils"
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
			utils.AssertEqual(0, len(context.Scope), t)
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
				utils.AssertEqual(value, context.GetVariable(key), t)
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
			utils.AssertErrorsEqual(constant.BuildCommandError, err, t)
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
			utils.AssertErrorsEqual(constant.RunCommandError, err, t)
			return
		case <-context.GetProcessingChannels().Success:
			t.Log("Should not got success result")
			t.Fail()
		}
	}
}
