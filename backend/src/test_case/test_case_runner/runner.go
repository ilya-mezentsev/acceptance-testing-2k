package test_case_runner

import (
	"interfaces"
	"models"
)

type Runner struct {
	transactions []interfaces.Transaction
}

func (r Runner) Run(processing models.TestsRun) {
	panic("implement me")
}

func (r *Runner) AddTransaction(transaction interfaces.Transaction) {
	r.transactions = append(r.transactions, transaction)
}
