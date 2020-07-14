package response_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"utils"
)

func Parse(response http.Response) (map[string]interface{}, error) {
	responseData := map[string]interface{}{}

	err := unmarshalRequestBody(response.Body, responseData)
	if err != nil {
		return nil, err
	}

	err = processValues(responseData)

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

func processValues(data map[string]interface{}) error {
	for key, value := range data {
		newValue, err := processValue(value)
		if err != nil {
			return err
		}

		data[key] = newValue
	}

	return nil
}

func processValue(value interface{}) (interface{}, error) {
	switch value.(type) {
	case int, int32, int64, uint, uint32:
		return fmt.Sprintf("%d", value), nil
	case float32, float64:
		return removeUselessZeros(fmt.Sprintf("%.5f", value)), nil
	case bool:
		return fmt.Sprintf("%v", value), nil
	default:
		return processComplexValue(value)
	}
}

func removeUselessZeros(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		currentSymbol := string(num[i])
		if currentSymbol == "0" {
			continue
		}

		if currentSymbol == "." {
			if currentSymbol == "0" {
				return num[:i-1]
			} else {
				return num[:i]
			}
		}
	}

	return num
}

func processComplexValue(value interface{}) (interface{}, error) {
	switch {
	case utils.IsGenericSlice(value):
		return getValueFromSlice(value.([]interface{}))
	case utils.IsGenericMap(value):
		return getValueFromMap(value.(map[string]interface{}))
	default:
		return nil, nil
	}
}

func getValueFromSlice(value []interface{}) ([]interface{}, error) {
	var values []interface{}
	for _, item := range value {
		newValue, err := processValue(item)
		if err != nil {
			return nil, err
		}

		values = append(values, newValue)
	}

	return values, nil
}

func getValueFromMap(value map[string]interface{}) (map[string]interface{}, error) {
	values := map[string]interface{}{}
	for key, val := range value {
		newValue, err := processValue(val)
		if err != nil {
			return nil, err
		}

		values[key] = newValue
	}

	return values, nil
}
