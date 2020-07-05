package parser

import (
	"test_case/parsers/errors"
	"test_case/parsers/transaction/data"
	"testing"
	"utils"
)

func TestParseSimpleTransactionWithArguments(t *testing.T) {
	var transactionData = data.SimpleTransactionData{}
	err := Parse(
		data.SimpleTransactionPattern,
		`CREATE USER {"hash": "some-hash", "userName": "Piter"}`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("CREATE", transactionData.GetCommand(), t)
	utils.AssertEqual("USER", transactionData.GetObject(), t)
	utils.AssertEqual(`{"hash": "some-hash", "userName": "Piter"}`, transactionData.GetArguments(), t)
}

func TestParseSimpleTransactionWithoutArguments(t *testing.T) {
	var transactionData = data.SimpleTransactionData{}
	err := Parse(
		data.SimpleTransactionPattern,
		`CREATE USER`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("CREATE", transactionData.GetCommand(), t)
	utils.AssertEqual("USER", transactionData.GetObject(), t)
	utils.AssertEqual(``, transactionData.GetArguments(), t)
}

func TestParseSimpleTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.SimpleTransactionPattern,
		``,
		&data.SimpleTransactionData{},
	)

	utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}

func TestParseAssignmentTransactionWithArguments(t *testing.T) {
	var transactionData = data.AssignmentTransactionData{}
	err := Parse(
		data.AssignmentTransactionPattern,
		`x = GET USER {"hash": "hash-1"}`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("x", transactionData.GetVariableName(), t)
	utils.AssertEqual("GET", transactionData.GetCommand(), t)
	utils.AssertEqual("USER", transactionData.GetObject(), t)
	utils.AssertEqual(`{"hash": "hash-1"}`, transactionData.GetArguments(), t)
}

func TestParseAssignmentTransactionWithoutArguments(t *testing.T) {
	var transactionData = data.AssignmentTransactionData{}
	err := Parse(
		data.AssignmentTransactionPattern,
		`x = GET USER`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("x", transactionData.GetVariableName(), t)
	utils.AssertEqual("GET", transactionData.GetCommand(), t)
	utils.AssertEqual("USER", transactionData.GetObject(), t)
	utils.AssertEqual(``, transactionData.GetArguments(), t)
}

func TestParseAssignmentTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.AssignmentTransactionPattern,
		``,
		&data.AssignmentTransactionData{},
	)

	utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}

func TestParseAssertionTransactionWithoutDataPath(t *testing.T) {
	var transactionData = data.AssertionTransactionData{}
	err := Parse(
		data.AssertionTransactionPattern,
		`ASSERT user EQUALS Ron`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("user", transactionData.GetVariableName(), t)
	utils.AssertEqual("Ron", transactionData.GetNewValue(), t)
	utils.AssertEqual("", transactionData.GetDataPath(), t)
}

func TestParseAssertionTransactionWithDataPath(t *testing.T) {
	var transactionData = data.AssertionTransactionData{}
	err := Parse(
		data.AssertionTransactionPattern,
		`ASSERT user.userName EQUALS Ron`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("user", transactionData.GetVariableName(), t)
	utils.AssertEqual("Ron", transactionData.GetNewValue(), t)
	utils.AssertEqual("userName", transactionData.GetDataPath(), t)
}

func TestParseAssertionTransactionWithLongDataPath(t *testing.T) {
	var transactionData = data.AssertionTransactionData{}
	err := Parse(
		data.AssertionTransactionPattern,
		`ASSERT user.data.name.value EQUALS Ron`,
		&transactionData,
	)

	utils.AssertNil(err, t)
	utils.AssertEqual("user", transactionData.GetVariableName(), t)
	utils.AssertEqual("Ron", transactionData.GetNewValue(), t)
	utils.AssertEqual("data.name.value", transactionData.GetDataPath(), t)
}

func TestParseAssertionTransactionInvalidTransactionFormat(t *testing.T) {
	err := Parse(
		data.AssertionTransactionPattern,
		`ASSERT user.data.name.value EQUALS`,
		&data.AssertionTransactionData{},
	)

	utils.AssertErrorsEqual(errors.InvalidTransactionFormat, err, t)
}
