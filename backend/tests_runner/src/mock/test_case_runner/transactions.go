package test_case_runner

import (
	"interfaces"
	"models"
	"test_case/errors"
)

type baseMockTransaction struct {
	previousCallContext interfaces.TestCaseContext
}

func (t *baseMockTransaction) Execute(context interfaces.TestCaseContext) models.TransactionError {
	t.previousCallContext = context
	return errors.EmptyTransactionError
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

func (t *MockErroredTransaction) Execute(context interfaces.TestCaseContext) models.TransactionError {
	t.previousCallContext = context
	return models.TransactionError{Code: SomeTransactionError.Error()}
}
