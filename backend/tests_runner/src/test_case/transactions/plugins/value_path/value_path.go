package value_path

import (
	"strconv"
	"strings"
	"type_utils"
)

func GetByPath(value interface{}, path string) (interface{}, error) {
	switch {
	case path == "":
		return nil, invalidPath
	case type_utils.IsGenericMap(value):
		return getValueFromMap(value, path)
	case type_utils.IsGenericSlice(value):
		return getValueFromSlice(value, path)
	default:
		return nil, CannotAccessValueByPath
	}
}

func getValueFromMap(value interface{}, path string) (interface{}, error) {
	pathComponents := strings.Split(path, ".")
	value = value.(map[string]interface{})[pathComponents[0]]

	if len(pathComponents) < 2 {
		return value, nil
	} else {
		return GetByPath(value, strings.Join(pathComponents[1:], "."))
	}
}

func getValueFromSlice(value interface{}, path string) (interface{}, error) {
	pathComponents := strings.Split(path, ".")
	supposedStringIndex := pathComponents[0]
	supposedIndex, err := strconv.Atoi(supposedStringIndex)
	sliceValue := value.([]interface{})
	if err != nil {
		return nil, invalidNumberForIndex
	} else if supposedIndex > len(sliceValue)-1 {
		return nil, indexOutOfBounds
	}

	supposedValue := sliceValue[supposedIndex]
	if len(pathComponents) < 2 {
		return supposedValue, nil
	} else {
		return GetByPath(supposedValue, strings.Join(pathComponents[1:], "."))
	}
}
