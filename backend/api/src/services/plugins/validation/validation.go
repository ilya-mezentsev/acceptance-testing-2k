package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func IsValid(data interface{}) bool {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		panic("passed struct is not pointer")
	}

	dataValue := reflect.ValueOf(data).Elem()
	if dataValue.Type().Kind() != reflect.Struct {
		panic(fmt.Errorf("cannot validate type: %s", dataValue.Type().Name()))
	}

	validationResults := validateStruct(dataValue)
	for _, res := range validationResults {
		if !res {
			return false
		}
	}

	return len(validationResults) > 0
}

func validateStruct(dataValue reflect.Value) []bool {
	var validationResults []bool
	for i := 0; i < dataValue.NumField(); i++ {
		fieldValue := dataValue.Field(i)

		switch fieldValue.Type().Kind() {
		case reflect.Struct:
			validationResults = append(validationResults, validateStruct(fieldValue)...)
		case reflect.Slice:
			validationResults = append(validationResults, validateSlice(fieldValue)...)
		default:
			validationResults = append(
				validationResults,
				validateField(dataValue.Type().Field(i), fieldValue.String()),
			)

			if fieldValue.Kind() == reflect.Int {
				validationResults = append(
					validationResults,
					validIntRanges(dataValue.Type().Field(i), fieldValue.Int()),
				)
			}
		}
	}

	return validationResults
}

func validateSlice(slice reflect.Value) []bool {
	var results []bool
	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		if item.Kind() != reflect.Struct {
			panic("slice argument is not struct")
		}

		results = append(results, validateStruct(item)...)
	}

	return results
}

func validateField(field reflect.StructField, value string) bool {
	validationRule, ok := field.Tag.Lookup("validation")
	if !ok {
		// we do not need to validate if validation tag is not provided
		return true
	}

	validationMethod, hasValidationRule := validationRuleToMethod[validationRule]
	return hasValidationRule && validationMethod(value)
}

func validIntRanges(field reflect.StructField, value int64) bool {
	rangeRule, ok := field.Tag.Lookup("range")
	if !ok {
		// we do not need to validate if validation tag is not provided
		return true
	}

	parsedRange := strings.Split(rangeRule, ",")
	intMin, errMin := strconv.Atoi(parsedRange[0])
	intMax, errMax := strconv.Atoi(parsedRange[1])

	if errMin != nil || errMax != nil {
		switch {
		case errMin != nil:
			panic(fmt.Errorf("invalid min range value: %v", parsedRange[0]))
		case errMax != nil:
			panic(fmt.Errorf("invalid max range value: %v", parsedRange[1]))
		}
	}

	return int64(intMin) <= value && value <= int64(intMax)
}
