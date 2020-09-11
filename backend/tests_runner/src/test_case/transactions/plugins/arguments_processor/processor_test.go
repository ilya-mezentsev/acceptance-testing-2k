package arguments_processor

import (
	"test_case/transactions/plugins/value_path"
	mockContext "test_runner_meta/mock/transaction/context"
	"test_utils"
	"testing"
)

var context = mockContext.Mock

func TestProcessSuccess(t *testing.T) {
	defer context.ClearScope()

	context.SetVariable("user", map[string]interface{}{
		"hash": "some-hash",
		"name": "Joe",
		"age":  15,
	})

	processedArguments, err := ReplaceTemplatesWithVariables(
		context,
		`{"hash": "${user.hash}", "name": "${user.name}", "age": ${user.age}}`,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(
		`{"hash": "some-hash", "name": "Joe", "age": 15}`,
		processedArguments,
		t,
	)
}

func TestProcessNoVariablesInArguments(t *testing.T) {
	processedArguments, err := ReplaceTemplatesWithVariables(
		context,
		`{"hash": "some-hash", "name": "Joe"}`,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(
		`{"hash": "some-hash", "name": "Joe"}`,
		processedArguments,
		t,
	)
}

func TestProcessErrorVariableIsNotDefined(t *testing.T) {
	_, err := ReplaceTemplatesWithVariables(
		context,
		`{"hash": "${user.hash}", "name": "${user.name}"}`,
	)

	test_utils.AssertErrorsEqual(value_path.CannotAccessValueByPath, err, t)
}
