package test_case

import (
	"test_case/parsers/errors"
	"test_utils"
	"testing"
)

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
	iterator, err := Parse(testCases)

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(iterator, t)
	test_utils.AssertEqual(1, len(iterator), t)
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
	iterator, err := Parse(testCases)

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(iterator, t)
	test_utils.AssertEqual(2, len(iterator), t)
}

func TestParser_ParseEmptyTestCases(t *testing.T) {
	_, err := Parse(``)

	test_utils.AssertErrorsEqual(errors.NoTestCases, err, t)
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
	iterator, _ := Parse(testCases)

	test_utils.AssertTrue(iterator[0].HasTransactions(), t)
}

func TestTestCaseIterator_HasTransactionsFalse(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
		END
	`
	iterator, _ := Parse(testCases)

	test_utils.AssertFalse(iterator[0].HasTransactions(), t)
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
	iterators, _ := Parse(testCases)
	iterator := iterators[0]

	for iterator.HasTransactions() {
		test_utils.AssertEqual(
			expectedTransactions[iterator.currentTransactionIndex-1],
			iterator.GetTestCaseTransaction(),
			t,
		)
	}
}
