package test_case

import (
	"parsers/errors"
	"testing"
	"utils"
)

var parser Parser

func TestParser_InitSuccess(t *testing.T) {
	testCases := `
		{
			// some comment (will be ignored)
			CREATE USER hash:'some-hash' userName:'Piter'
	
			user = GET USER hash:'some-hash'
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		}
	`
	err := parser.Init(testCases)

	utils.AssertNil(err, t)
}

func TestParser_InitCoupleTestCases(t *testing.T) {
	testCases := `
		{
			// some comment (will be ignored)
			CREATE USER hash:'some-hash' userName:'Piter'
	
			user = GET USER hash:'some-hash'
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		}
		{
			UPDATE USER hash:'hash-1' userName:'Ron'
	
			user = GET USER hash:'hash-1'
	
			ASSERT user.userName EQUALS 'Ron'
		}
	`
	_ = parser.Init(testCases)

	utils.AssertEqual(2, len(parser.testCases), t)
	testCaseTransactions := parser.NextTransactions()
	for ; !parser.Done(); testCaseTransactions = parser.NextTransactions() {
		for _, transaction := range testCaseTransactions {
			utils.AssertNotEqual("//", transaction[:2], t)
			utils.AssertNotEqual("", transaction, t)
		}
	}
}

func TestParser_InitError(t *testing.T) {
	err := parser.Init(``)

	utils.AssertErrorsEqual(errors.NoTestCases, err, t)
}

func TestParser_Next(t *testing.T) {
	testCases := `
		{
			// some comment (will be ignored)
			CREATE USER hash:'some-hash' userName:'Piter'
	
			user = GET USER hash:'some-hash'
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		}
	`
	_ = parser.Init(testCases)
	testCase := parser.NextTransactions()

	utils.AssertEqual(
		`CREATE USER hash:'some-hash' userName:'Piter'`,
		testCase[0],
		t,
	)
	utils.AssertEqual(
		`user = GET USER hash:'some-hash'`,
		testCase[1],
		t,
	)
	utils.AssertEqual(
		`ASSERT user.hash EQUALS 'some-hash'`,
		testCase[2],
		t,
	)
	utils.AssertEqual(
		`ASSERT user.userName EQUALS 'Piter'`,
		testCase[3],
		t,
	)
}

func TestParser_Done(t *testing.T) {
	testCases := `
		{
			// some comment (will be ignored)
			CREATE USER hash:'some-hash' userName:'Piter'
	
			user = GET USER hash:'some-hash'
	
			ASSERT user.hash EQUALS 'some-hash'
			ASSERT user.userName EQUALS 'Piter'
		}
	`
	_ = parser.Init(testCases)

	utils.AssertFalse(parser.Done(), t)
	_ = parser.NextTransactions()
	utils.AssertTrue(parser.Done(), t)
}
