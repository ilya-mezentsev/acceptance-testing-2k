package runner

import (
	"test_case/errors"
	"test_runner_meta/interfaces"
	"test_runner_meta/models"
)

type Runner struct {
	context      interfaces.TestCaseContext
	transactions []interfaces.Transaction
}

func (r *Runner) Run(processing models.TestsRun) {
	if len(r.transactions) < 1 {
		processing.Error <- models.TransactionError{Code: errors.NoTransactionsInTestCase.Error()}
		return
	}

	r.initTestCaseContext()
	for _, transaction := range r.transactions {
		transactionError := transaction.Execute(r.context)
		if transactionError != errors.EmptyTransactionError {
			processing.Error <- transactionError
			return
		}
	}

	processing.Success <- true
}

func (r *Runner) initTestCaseContext() {
	r.context = &Context{
		Scope: map[string]interface{}{},
	}
}

func (r *Runner) AddTransaction(transaction interfaces.Transaction) {
	r.transactions = append(r.transactions, transaction)
}
