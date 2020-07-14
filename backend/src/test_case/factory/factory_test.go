package factory

import (
	"test_case/errors"
	parseErrors "test_case/parsers/errors"
	"testing"
	"utils"
)

var factory = New(nil)

func TestFactory_GetAllSuccess(t *testing.T) {
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
	runners, err := factory.GetAll(testCases)

	utils.AssertNil(err, t)
	utils.AssertNotNil(runners, t)
}

func TestFactory_GetAllEmptyTestCases(t *testing.T) {
	_, err := factory.GetAll(``)

	utils.AssertErrorsEqual(parseErrors.NoTestCases, err, t)
}

func TestFactory_GetAllUnknownTransactionType(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
			blah
		END
	`
	_, err := factory.GetAll(testCases)

	utils.AssertErrorsEqual(errors.UnknownTransactionType, err, t)
}

func TestFactory_GetAssertionTransactionError(t *testing.T) {
	_, err := factory.(Factory).getAssertionTransaction(``)

	utils.AssertErrorsEqual(parseErrors.InvalidTransactionFormat, err, t)
}

func TestFactory_GetAssignmentTransactionError(t *testing.T) {
	_, err := factory.(Factory).getAssignmentTransaction(``)

	utils.AssertErrorsEqual(parseErrors.InvalidTransactionFormat, err, t)
}

func TestFactory_GetSimpleTransactionError(t *testing.T) {
	_, err := factory.(Factory).getSimpleTransaction(``)

	utils.AssertErrorsEqual(parseErrors.InvalidTransactionFormat, err, t)
}
