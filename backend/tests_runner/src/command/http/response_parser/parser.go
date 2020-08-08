package response_parser

import (
	"command/http/errors"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"plugins/logger"
	"strings"
	"type_utils"
)

func Parse(response http.Response) (map[string]interface{}, error) {
	responseData := map[string]interface{}{}
	err := unmarshalRequestBody(response.Body, responseData)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to parse JSON from response body: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"response_body": response.Body,
			},
		}, logger.Error)

		return nil, errors.NoJSONInResponse
	}

	for key, value := range responseData {
		responseData[key] = processValue(value)
	}

	return responseData, err
}

func unmarshalRequestBody(closer io.ReadCloser, dest map[string]interface{}) error {
	responseBody, err := ioutil.ReadAll(closer)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, &dest)
	if err != nil {
		return err
	}

	return nil
}

func processValue(value interface{}) interface{} {
	switch value.(type) {
	case string:
		return value
	case float32, float64:
		return removeUselessZeros(fmt.Sprintf("%.5f", value))
	case bool:
		return fmt.Sprintf("%v", value)
	}

	switch {
	case type_utils.IsGenericSlice(value):
		return getValueFromSlice(value.([]interface{}))
	case type_utils.IsGenericMap(value):
		return getValueFromMap(value.(map[string]interface{}))
	}

	logger.WithFields(logger.Fields{
		MessageTemplate: "Unable to process value.",
		Optional: map[string]interface{}{
			"type":  fmt.Sprintf("%T", value),
			"value": fmt.Sprintf("%v", value),
		},
	}, logger.Warning)
	return nil
}

func removeUselessZeros(num string) string {
	splatted := strings.Split(num, ".")
	whole, fractional := splatted[0], splatted[1]
	var newFractional string

	for i := len(fractional) - 1; i >= 0; i-- {
		currentSymbol := string(fractional[i])

		if currentSymbol != "0" {
			newFractional = fractional[:i+1]
		}
	}

	if newFractional == "" {
		return whole
	} else {
		return strings.Join([]string{whole, newFractional}, ".")
	}
}

func getValueFromSlice(value []interface{}) []interface{} {
	var values []interface{}
	for _, item := range value {
		values = append(values, processValue(item))
	}

	return values
}

func getValueFromMap(value map[string]interface{}) map[string]interface{} {
	values := map[string]interface{}{}
	for key, val := range value {
		values[key] = processValue(val)
	}

	return values
}
