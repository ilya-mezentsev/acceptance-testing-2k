package runner

import (
	mockTestCaseRunner "mock/test_case_runner"
	"models"
	"test_case/errors"
	"test_utils"
	"testing"
)

func TestRunner_RunOneSimpleTransaction(t *testing.T) {
	var runner Runner
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}
	transaction := mockTestCaseRunner.MockTransaction{}

	runner.AddTransaction(&transaction)

	go runner.Run(processing)

	for {
		select {
		case <-processing.Success:
			test_utils.AssertTrue(transaction.CalledWith(runner.context), t)
			return
		case err := <-processing.Error:
			t.Log(err)
			t.Fail()
		}
	}
}

func TestRunner_RunFewSimpleTransactions(t *testing.T) {
	var runner Runner
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}
	transaction1 := mockTestCaseRunner.MockTransaction{}
	transaction2 := mockTestCaseRunner.MockTransaction{}

	runner.AddTransaction(&transaction1)
	runner.AddTransaction(&transaction2)

	go runner.Run(processing)

	for {
		select {
		case <-processing.Success:
			test_utils.AssertTrue(transaction1.CalledWith(runner.context), t)
			test_utils.AssertTrue(transaction2.CalledWith(runner.context), t)
			return
		case err := <-processing.Error:
			t.Log(err)
			t.Fail()
		}
	}
}

func TestRunner_RunOneErroredTransaction(t *testing.T) {
	var runner Runner
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}
	transaction := mockTestCaseRunner.ErroredMockTransaction

	runner.AddTransaction(&transaction)

	go runner.Run(processing)

	for {
		select {
		case <-processing.Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-processing.Error:
			test_utils.AssertEqual(mockTestCaseRunner.SomeTransactionError.Error(), err.Code, t)
			test_utils.AssertTrue(transaction.CalledWith(runner.context), t)
			return
		}
	}
}

func TestRunner_RunFirstSimpleThenErroredTransactions(t *testing.T) {
	var runner Runner
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}
	simpleTransaction := mockTestCaseRunner.SimpleMockTransaction
	erroredTransaction := mockTestCaseRunner.ErroredMockTransaction

	runner.AddTransaction(&simpleTransaction)
	runner.AddTransaction(&erroredTransaction)

	go runner.Run(processing)

	for {
		select {
		case <-processing.Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-processing.Error:
			test_utils.AssertEqual(mockTestCaseRunner.SomeTransactionError.Error(), err.Code, t)
			test_utils.AssertTrue(simpleTransaction.CalledWith(runner.context), t)
			test_utils.AssertTrue(erroredTransaction.CalledWith(runner.context), t)
			return
		}
	}
}

func TestRunner_RunFirstErroredThenSimpleTransactions(t *testing.T) {
	var runner Runner
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}
	simpleTransaction := mockTestCaseRunner.SimpleMockTransaction
	erroredTransaction := mockTestCaseRunner.ErroredMockTransaction

	runner.AddTransaction(&erroredTransaction)
	runner.AddTransaction(&simpleTransaction)

	go runner.Run(processing)

	for {
		select {
		case <-processing.Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-processing.Error:
			test_utils.AssertEqual(mockTestCaseRunner.SomeTransactionError.Error(), err.Code, t)
			test_utils.AssertFalse(simpleTransaction.CalledWith(runner.context), t)
			test_utils.AssertTrue(erroredTransaction.CalledWith(runner.context), t)
			return
		}
	}
}

func TestRunner_RunNoTransactions(t *testing.T) {
	var runner Runner
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}

	go runner.Run(processing)

	for {
		select {
		case <-processing.Success:
			t.Log("Should not got success result")
			t.Fail()
			return
		case err := <-processing.Error:
			test_utils.AssertEqual(errors.NoTransactionsInTestCase.Error(), err.Code, t)
			return
		}
	}
}

func TestContext_GetExistsVariable(t *testing.T) {
	context := Context{
		Scope: map[string]interface{}{},
	}
	context.SetVariable("x", "10")

	test_utils.AssertEqual("10", context.GetVariable("x"), t)
}

func TestContext_GetNotExistsVariable(t *testing.T) {
	context := Context{
		Scope: map[string]interface{}{},
	}

	test_utils.AssertNil(context.GetVariable("x"), t)
}
