package arguments_builder

import (
	"command/http/errors"
	"encoding/json"
	"fmt"
	"logger"
	"regexp"
	"strings"
)

var slashSeparatedPattern = regexp.MustCompile(
	`^(/?[\w-]+?/?)+$`,
)

type arguments struct {
	data string
}

func (a arguments) Value() string {
	return a.data
}

func (a arguments) IsSlashSeparated() bool {
	return slashSeparatedPattern.MatchString(a.data)
}

func (a arguments) AmpersandSeparated() (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(a.data), &data)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to parse arguments string: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"data": a.data,
			},
		}, logger.Error)

		return "", errors.NoJSONInArguments
	}

	var ampersandSeparated []string
	for key, value := range data {
		switch value.(type) {
		case string, bool, float32, float64:
			ampersandSeparated = append(
				ampersandSeparated,
				fmt.Sprintf("%s=%v", key, value),
			)
		default:
			jsonEncodedValue, _ := json.Marshal(value)

			ampersandSeparated = append(
				ampersandSeparated,
				fmt.Sprintf("%s=%s", key, jsonEncodedValue),
			)
		}
	}

	return strings.Join(ampersandSeparated, "&"), nil
}
