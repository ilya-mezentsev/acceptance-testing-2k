package factory

import (
	"test_case/errors"
	"test_case/parsers/test_case"
	"test_case/parsers/transaction/data"
	"test_case/parsers/transaction/parser"
	"test_case/runner"
	"test_case/transactions/assertion"
	"test_case/transactions/assignment"
	"test_runner_meta/interfaces"
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
			transaction, err := f.getTransaction(
				testCasesIterator.GetTestCaseTransaction(),
				testCasesIterator.GetTestCaseText(),
			)
			if err != nil {
				return nil, err
			}

			testCaseRunner.AddTransaction(transaction)
		}

		testCaseRunners = append(testCaseRunners, &testCaseRunner)
	}

	return testCaseRunners, nil
}

func (f Factory) getTransaction(
	transactionText,
	testCaseText string,
) (interfaces.Transaction, error) {
	switch {
	case data.IsAssertion(transactionText):
		return f.getAssertionTransaction(transactionText, testCaseText)
	case data.IsFalseAssertion(transactionText):
		return f.getFalseAssertionTransaction(transactionText, testCaseText)
	case data.IsTrueAssertion(transactionText):
		return f.getTrueAssertionTransaction(transactionText, testCaseText)
	case data.IsAssignment(transactionText):
		return f.getAssignmentTransaction(transactionText, testCaseText)
	default:
		return nil, errors.UnknownTransactionType
	}
}

func (f Factory) getAssertionTransaction(
	transactionText,
	testCaseText string,
) (interfaces.Transaction, error) {
	var assertionTransactionData data.AssertionTransactionData
	err := parser.Parse(
		data.AssertionTransactionPattern,
		transactionText,
		&assertionTransactionData,
	)
	if err != nil {
		return nil, err
	}

	assertionTransactionData.SetTestCaseText(testCaseText)
	return assertion.New(&assertionTransactionData), nil
}

func (f Factory) getFalseAssertionTransaction(
	transactionText,
	testCaseText string,
) (interfaces.Transaction, error) {
	var falseAssertionTransactionData data.FalseAssertionTransactionData
	err := parser.Parse(
		data.FalseAssertionTransactionPattern,
		transactionText,
		&falseAssertionTransactionData,
	)
	if err != nil {
		return nil, err
	}

	falseAssertionTransactionData.SetTestCaseText(testCaseText)
	return assertion.New(&falseAssertionTransactionData), nil
}

func (f Factory) getTrueAssertionTransaction(
	transactionText,
	testCaseText string,
) (interfaces.Transaction, error) {
	var trueAssertionTransactionData data.TrueAssertionTransactionData
	err := parser.Parse(
		data.TrueAssertionTransactionPattern,
		transactionText,
		&trueAssertionTransactionData,
	)
	if err != nil {
		return nil, err
	}

	trueAssertionTransactionData.SetTestCaseText(testCaseText)
	return assertion.New(&trueAssertionTransactionData), nil
}

func (f Factory) getAssignmentTransaction(
	transactionText,
	testCaseText string,
) (interfaces.Transaction, error) {
	var assignmentTransactionData data.AssignmentTransactionData
	err := parser.Parse(
		data.AssignmentTransactionPattern,
		transactionText,
		&assignmentTransactionData,
	)
	if err != nil {
		return nil, err
	}

	assignmentTransactionData.SetTestCaseText(testCaseText)
	return assignment.New(f.commandBuilder, &assignmentTransactionData), nil
}
