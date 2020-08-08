package test_case_runner

import (
	"test_case/errors"
	"test_runner_meta/interfaces"
	"test_runner_meta/models"
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
