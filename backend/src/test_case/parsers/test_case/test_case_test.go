package test_case

import (
	"test_case/parsers/errors"
	"testing"
	"utils"
)

var parser Parser

func TestParser_ParseOneTestCase(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
			CREATE USER {"hash": "some-hash", "userName": "Piter"}
	
			user = GET USER {"hash": "some-hash"}
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		END
	`
	iterator, err := parser.Parse(testCases)

	utils.AssertNil(err, t)
	utils.AssertNotNil(iterator, t)
	utils.AssertEqual(1, len(iterator), t)
}

func TestParser_ParseTwoTestCases(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
			CREATE USER {"hash": "some-hash", "userName": "Piter"}
	
			user = GET USER {"hash": "some-hash"}
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		END
		BEGIN
			user = GET USER {"hash": "some-hash"}
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		END
	`
	iterator, err := parser.Parse(testCases)

	utils.AssertNil(err, t)
	utils.AssertNotNil(iterator, t)
	utils.AssertEqual(2, len(iterator), t)
}

func TestParser_ParseEmptyTestCases(t *testing.T) {
	_, err := parser.Parse(``)

	utils.AssertErrorsEqual(errors.NoTestCases, err, t)
}

func TestTestCaseIterator_HasTransactionsTrue(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
			CREATE USER {"hash": "some-hash", "userName": "Piter"}
	
			user = GET USER {"hash": "some-hash"}
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		END
	`
	iterator, _ := parser.Parse(testCases)

	utils.AssertTrue(iterator[0].HasTransactions(), t)
}

func TestTestCaseIterator_HasTransactionsFalse(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
		END
	`
	iterator, _ := parser.Parse(testCases)

	utils.AssertFalse(iterator[0].HasTransactions(), t)
}

func TestTestCaseIterator_GetTestCaseTransactions(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
			CREATE USER {"hash": "some-hash", "userName": "Piter"}
	
			user = GET USER {"hash": "some-hash"}
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		END
	`
	expectedTransactions := []string{
		`CREATE USER {"hash": "some-hash", "userName": "Piter"}`,
		`user = GET USER {"hash": "some-hash"}`,
		`ASSERT user.hash EQUALS 'some-hash'`,
		`ASSERT user.userName EQUALS 'Piter'`,
	}
	iterators, _ := parser.Parse(testCases)
	iterator := iterators[0].(*TestCaseTransactionsIterator)

	for iterator.HasTransactions() {
		utils.AssertEqual(
			expectedTransactions[iterator.currentTransactionIndex-1],
			iterator.GetTestCaseTransaction(),
			t,
		)
	}
}
