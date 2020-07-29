package utils

func IsGenericSlice(value interface{}) bool {
	_, ok := value.([]interface{})

	return ok
}

func IsGenericMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})

	return ok
}
