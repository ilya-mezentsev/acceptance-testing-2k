package assertion

import (
	"strconv"
	"strings"
)

func getValueByPath(value interface{}, path string) (interface{}, error) {
	switch {
	case path == "":
		return nil, invalidPath
	case isMap(value):
		return getValueFromMap(value, path)
	case isSlice(value):
		return getValueFromSlice(value, path)
	default:
		return nil, cannotAccessValueByPath
	}
}

func isMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})

	return ok
}

func isSlice(value interface{}) bool {
	_, ok := value.([]interface{})

	return ok
}

func getValueFromMap(value interface{}, path string) (interface{}, error) {
	pathComponents := strings.Split(path, ".")
	value = value.(map[string]interface{})[pathComponents[0]]

	if len(pathComponents) < 2 {
		return value, nil
	} else {
		return getValueByPath(value, strings.Join(pathComponents[1:], "."))
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
		return getValueByPath(supposedValue, strings.Join(pathComponents[1:], "."))
	}
}
