package factory

import (
	"interfaces"
	"test_case/errors"
	"test_case/parsers/test_case"
	"test_case/parsers/transaction/data"
	"test_case/parsers/transaction/parser"
	"test_case/runner"
	"test_case/transactions/assertion"
	"test_case/transactions/assignment"
	"test_case/transactions/simple"
)

type Factory struct {
	commandBuilder interfaces.CommandBuilder
}

func New(commandBuilder interfaces.CommandBuilder) interfaces.TestCaseFactory {
	return Factory{commandBuilder}
}

func (f Factory) GetAll(testCasesData string) ([]interfaces.TestCaseRunner, error) {
	testCasesIterators, err := test_case.Parse(testCasesData)
	if err != nil {
		return nil, err
	}

	var testCaseRunners []interfaces.TestCaseRunner
	for _, testCasesIterator := range testCasesIterators {
		testCaseRunner := runner.Runner{}
		for testCasesIterator.HasTransactions() {
			transactionData := testCasesIterator.GetTestCaseTransaction()
			transaction, err := f.getTransaction(transactionData)
			if err != nil {
				return nil, err
			}

			testCaseRunner.AddTransaction(transaction)
		}

		testCaseRunners = append(testCaseRunners, &testCaseRunner)
	}

	return testCaseRunners, nil
}

func (f Factory) getTransaction(transactionData string) (interfaces.Transaction, error) {
	switch {
	case data.IsAssertion(transactionData):
		return f.getAssertionTransaction(transactionData)
	case data.IsAssignment(transactionData):
		return f.getAssignmentTransaction(transactionData)
	case data.IsSimple(transactionData):
		return f.getSimpleTransaction(transactionData)
	default:
		return nil, errors.UnknownTransactionType
	}
}

func (f Factory) getAssertionTransaction(transactionData string) (interfaces.Transaction, error) {
	var assertionTransactionData data.AssertionTransactionData
	err := parser.Parse(
		data.AssertionTransactionPattern,
		transactionData,
		&assertionTransactionData,
	)
	if err != nil {
		return nil, err
	}

	return assertion.New(&assertionTransactionData), nil
}

func (f Factory) getAssignmentTransaction(transactionData string) (interfaces.Transaction, error) {
	var assignmentTransactionData data.AssignmentTransactionData
	err := parser.Parse(
		data.AssignmentTransactionPattern,
		transactionData,
		&assignmentTransactionData,
	)
	if err != nil {
		return nil, err
	}

	return assignment.New(f.commandBuilder, &assignmentTransactionData), nil
}

func (f Factory) getSimpleTransaction(transactionData string) (interfaces.Transaction, error) {
	var simpleTransactionData data.SimpleTransactionData
	err := parser.Parse(
		data.SimpleTransactionPattern,
		transactionData,
		&simpleTransactionData,
	)
	if err != nil {
		return nil, err
	}

	return simple.New(f.commandBuilder, &simpleTransactionData), nil
}
