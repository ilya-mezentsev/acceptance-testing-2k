package assertion

import "strings"

func getValueByPath(value interface{}, path string) (interface{}, error) {
	switch {
	case !isMap(value):
		return nil, cannotAccessValueByPath
	case path == "":
		return nil, invalidPath
	}

	pathComponents := strings.Split(path, ".")
	value = value.(map[string]interface{})[pathComponents[0]]
	if len(pathComponents) < 2 {
		return value, nil
	} else {
		return getValueByPath(value, strings.Join(pathComponents[1:], "."))
	}
}

func isMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})

	return ok
}
