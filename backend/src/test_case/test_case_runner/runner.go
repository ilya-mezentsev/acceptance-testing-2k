package test_case_runner

import (
	"interfaces"
	"models"
	"test_case/errors"
)

type Runner struct {
	context      interfaces.TestCaseContext
	transactions []interfaces.Transaction
}

func (r *Runner) Run(processing models.TestsRun) {
	if len(r.transactions) < 1 {
		processing.Error <- errors.NoTransactionsInTestCase
		return
	}

	r.initTestCaseContext()
	for _, transaction := range r.transactions {
		go transaction.Execute(r.context)

		for {
			select {
			case <-r.context.GetProcessingChannels().Success:
				goto FinishTransaction
			case err := <-r.context.GetProcessingChannels().Error:
				processing.Error <- err
				return
			}

		FinishTransaction:
			break
		}
	}

	processing.Success <- true
}

func (r *Runner) initTestCaseContext() {
	r.context = &Context{
		Scope: map[string]interface{}{},
		ProcessingChannels: models.TestsRun{
			Success: make(chan bool),
			Error:   make(chan error),
		},
	}
}

func (r *Runner) AddTransaction(transaction interfaces.Transaction) {
	r.transactions = append(r.transactions, transaction)
}
