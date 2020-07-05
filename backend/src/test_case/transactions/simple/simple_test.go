package simple

import (
	"mock/transaction/simple"
	"testing"
	"utils"
)

func TestTransaction_ExecuteSuccess(t *testing.T) {
	transaction := New(
		simple.MockCommandBuilder{},
		&simple.MockData,
	)

	go transaction.Execute(simple.Context)

	for {
		select {
		case err := <-simple.Context.GetProcessingChannels().Error:
			t.Log(err)
			t.Fail()
		case res := <-simple.Context.GetProcessingChannels().Success:
			utils.AssertTrue(res, t)
			return
		}
	}
}

func TestTransaction_ExecuteBuildCommandError(t *testing.T) {
	transaction := New(
		simple.MockCommandBuilderWithError{},
		&simple.MockData,
	)

	go transaction.Execute(simple.Context)

	for {
		select {
		case err := <-simple.Context.GetProcessingChannels().Error:
			utils.AssertErrorsEqual(simple.BuildCommandError, err, t)
			return
		}
	}
}

func TestTransaction_ExecuteCommandRunError(t *testing.T) {
	transaction := New(
		simple.MockErroredCommandBuilder{},
		&simple.MockData,
	)

	go transaction.Execute(simple.Context)

	for {
		select {
		case err := <-simple.Context.GetProcessingChannels().Error:
			utils.AssertErrorsEqual(simple.RunCommandError, err, t)
			return
		}
	}
}
