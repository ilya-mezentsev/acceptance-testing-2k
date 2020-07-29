package test_case_runner

import (
	"interfaces"
	"models"
)

type baseMockTransaction struct {
	previousCallContext interfaces.TestCaseContext
}

func (t *baseMockTransaction) Execute(context interfaces.TestCaseContext) {
	t.previousCallContext = context
	context.GetProcessingChannels().Success <- true
}

func (t baseMockTransaction) CalledWith(context interfaces.TestCaseContext) bool {
	return t.previousCallContext == context
}

type MockTransaction struct {
	baseMockTransaction
}

type MockErroredTransaction struct {
	baseMockTransaction
}

func (t *MockErroredTransaction) Execute(context interfaces.TestCaseContext) {
	t.previousCallContext = context
	context.GetProcessingChannels().Error <- models.TransactionError{Code: SomeTransactionError.Error()}
}
