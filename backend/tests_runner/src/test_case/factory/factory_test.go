package factory

import (
	"test_case/errors"
	parseErrors "test_case/parsers/errors"
	"test_utils"
	"testing"
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

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(runners, t)
}

func TestFactory_GetAllEmptyTestCases(t *testing.T) {
	_, err := factory.GetAll(``)

	test_utils.AssertErrorsEqual(parseErrors.NoTestCases, err, t)
}

func TestFactory_GetAllUnknownTransactionType(t *testing.T) {
	testCases := `
		BEGIN
			// some comment (will be ignored)
			blah
		END
	`
	_, err := factory.GetAll(testCases)

	test_utils.AssertErrorsEqual(errors.UnknownTransactionType, err, t)
}

func TestFactory_GetAssertionTransactionError(t *testing.T) {
	_, err := factory.(Factory).getAssertionTransaction(``, ``)

	test_utils.AssertErrorsEqual(parseErrors.InvalidTransactionFormat, err, t)
}

func TestFactory_GetAssignmentTransactionError(t *testing.T) {
	_, err := factory.(Factory).getAssignmentTransaction(``, ``)

	test_utils.AssertErrorsEqual(parseErrors.InvalidTransactionFormat, err, t)
}

func TestFactory_GetSimpleTransactionError(t *testing.T) {
	_, err := factory.(Factory).getSimpleTransaction(``, ``)

	test_utils.AssertErrorsEqual(parseErrors.InvalidTransactionFormat, err, t)
}
