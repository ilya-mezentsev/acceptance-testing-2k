package arguments_processor

import (
	"fmt"
	"regexp"
	"strings"
	"test_case/transactions/plugins/value_path"
	"test_runner_meta/interfaces"
)

var (
	variableReferencePattern = regexp.MustCompile(
		`(?:\${(?P<variableName>[a-zA-Z0-9_]+)(?:\.(?P<dataPath>[a-zA-Z0-9_.]+?))?})+`,
	)
)

// Function for replace ${...} pattern with variable from context
func ReplaceTemplatesWithVariables(
	context interfaces.TestCaseContext,
	arguments string,
) (string, error) {
	matches := variableReferencePattern.FindAllStringSubmatch(arguments, -1)

	for _, match := range matches {
		var variableName, dataPath string
		for i, groupName := range variableReferencePattern.SubexpNames() {
			if i == 0 || groupName == "" {
				continue
			}

			if groupName == "variableName" {
				variableName = match[i]
			} else if groupName == "dataPath" {
				dataPath = match[i]
			}
		}

		value, err := value_path.GetByPath(
			context.GetVariable(variableName),
			dataPath,
		)
		if err != nil {
			return "", err
		}

		arguments = strings.ReplaceAll(
			arguments,
			fmt.Sprintf("${%s.%s}", variableName, dataPath),
			fmt.Sprintf("%v", value),
		)
	}

	return arguments, nil
}
