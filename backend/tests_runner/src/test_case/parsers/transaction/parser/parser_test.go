package parser

import (
	"test_case/parsers/errors"
	"test_case/parsers/transaction/data"
	"test_utils"
	"testing"
)

func TestParseAssignmentTransactionWithJSONArguments(t *testing.T) {
	var transactionData data.AssignmentTransactionData
	transactionText := `x = GET USER {"hash": "hash-1"}`
	err := Parse(
		data.AssignmentTransactionPattern,
		transactionText,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("x", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("GET", transactionData.GetCommand(), t)
	test_utils.AssertEqual("USER", transactionData.GetObject(), t)
	test_utils.AssertEqual(`{"hash": "hash-1"}`, transactionData.GetArguments(), t)
	test_utils.AssertEqual(transactionText, transactionData.GetTransactionText(), t)
}

func TestParseAssignmentTransactionWithSlashSeparatedArguments(t *testing.T) {
	var transactionData data.AssignmentTransactionData
	transactionText := `x = GET USER hash-1/nickname`
	err := Parse(
		data.AssignmentTransactionPattern,
		transactionText,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("x", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("GET", transactionData.GetCommand(), t)
	test_utils.AssertEqual("USER", transactionData.GetObject(), t)
	test_utils.AssertEqual(`hash-1/nickname`, transactionData.GetArguments(), t)
	test_utils.AssertEqual(transactionText, transactionData.GetTransactionText(), t)
}

func TestParseAssignmentTransactionWithoutArguments(t *testing.T) {
	var transactionData data.AssignmentTransactionData
	transactionText := `x = GET USER`
	err := Parse(
		data.AssignmentTransactionPattern,
		transactionText,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("x", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("GET", transactionData.GetCommand(), t)
	test_utils.AssertEqual("USER", transactionData.GetObject(), t)
	test_utils.AssertEqual(``, transactionData.GetArguments(), t)
	test_utils.AssertEqual(transactionText, transactionData.GetTransactionText(), t)
}

func TestParseAssignmentTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.AssignmentTransactionPattern,
		``,
		&data.AssignmentTransactionData{},
	)

	test_utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}

func TestParseAssertionTransactionWithoutDataPath(t *testing.T) {
	var transactionData data.AssertionTransactionData
	transactionText := `ASSERT user EQUALS Ron`
	err := Parse(
		data.AssertionTransactionPattern,
		transactionText,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("user", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("Ron", transactionData.GetNewValue(), t)
	test_utils.AssertEqual("", transactionData.GetDataPath(), t)
	test_utils.AssertEqual(transactionText, transactionData.GetTransactionText(), t)
}

func TestParseAssertionTransactionWithDataPath(t *testing.T) {
	var transactionData data.AssertionTransactionData
	err := Parse(
		data.AssertionTransactionPattern,
		`ASSERT user.userName EQUALS Ron`,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("user", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("Ron", transactionData.GetNewValue(), t)
	test_utils.AssertEqual("userName", transactionData.GetDataPath(), t)
}

func TestParseAssertionTransactionWithLongDataPath(t *testing.T) {
	var transactionData data.AssertionTransactionData
	transactionText := `ASSERT user.data.name.value EQUALS Ron`
	err := Parse(
		data.AssertionTransactionPattern,
		transactionText,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("user", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("Ron", transactionData.GetNewValue(), t)
	test_utils.AssertEqual("data.name.value", transactionData.GetDataPath(), t)
	test_utils.AssertEqual(transactionText, transactionData.GetTransactionText(), t)
}

func TestParseAssertionTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.AssertionTransactionPattern,
		`ASSERT user.data.name.value EQUALS`,
		&data.AssertionTransactionData{},
	)

	test_utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}

func TestParseFalseAssertionTransaction(t *testing.T) {
	var transactionData data.FalseAssertionTransactionData
	err := Parse(
		data.FalseAssertionTransactionPattern,
		`ASSERT_FALSE user.data.exists`,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("user", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("false", transactionData.GetNewValue(), t)
	test_utils.AssertEqual("data.exists", transactionData.GetDataPath(), t)
	test_utils.AssertEqual(
		`ASSERT_FALSE user.data.exists`,
		transactionData.GetTransactionText(),
		t,
	)
}

func TestParseFalseAssertionTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.FalseAssertionTransactionPattern,
		`ASSERT_FALSE `,
		&data.FalseAssertionTransactionData{},
	)

	test_utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}

func TestParseTrueAssertionTransaction(t *testing.T) {
	var transactionData data.TrueAssertionTransactionData
	err := Parse(
		data.TrueAssertionTransactionPattern,
		`ASSERT_TRUE user.data.exists`,
		&transactionData,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("user", transactionData.GetVariableName(), t)
	test_utils.AssertEqual("true", transactionData.GetNewValue(), t)
	test_utils.AssertEqual("data.exists", transactionData.GetDataPath(), t)
	test_utils.AssertEqual(
		`ASSERT_TRUE user.data.exists`,
		transactionData.GetTransactionText(),
		t,
	)
}

func TestParseTrueAssertionTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.FalseAssertionTransactionPattern,
		`ASSERT_TRUE `,
		&data.TrueAssertionTransactionData{},
	)

	test_utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}
