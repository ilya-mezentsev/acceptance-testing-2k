package value_path

import (
	"test_utils"
	"testing"
)

func TestTransaction_GetByPathSingleKey(t *testing.T) {
	value, err := GetByPath(map[string]interface{}{
		"x": 10,
	}, "x")

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(10, value.(int), t)
}

func TestTransaction_GetByPathDotSeparated(t *testing.T) {
	value, err := GetByPath(map[string]interface{}{
		"x": map[string]interface{}{
			"y": 10,
		},
	}, "x.y")

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(10, value.(int), t)
}

func TestTransaction_GetByPathArray(t *testing.T) {
	value, err := GetByPath(map[string]interface{}{
		"x": []interface{}{1},
	}, "x.0")

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(1, value.(int), t)
}

func TestTransaction_GetByPathArrayWithMap(t *testing.T) {
	value, err := GetByPath(map[string]interface{}{
		"x": []interface{}{map[string]interface{}{
			"y": 1,
		}},
	}, "x.0.y")

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(1, value.(int), t)
}

func TestTransaction_GetByPathArrayIndexOutOfBounds(t *testing.T) {
	_, err := GetByPath(map[string]interface{}{
		"x": []interface{}{1},
	}, "x.1")

	test_utils.AssertErrorsEqual(indexOutOfBounds, err, t)
}

func TestTransaction_GetByPathArrayInvalidIndex(t *testing.T) {
	_, err := GetByPath(map[string]interface{}{
		"x": []interface{}{1},
	}, "x.a")

	test_utils.AssertErrorsEqual(invalidNumberForIndex, err, t)
}

func TestTransaction_GetByPathInvalidPath(t *testing.T) {
	_, err := GetByPath(map[string]interface{}{
		"x": 1,
	}, "")

	test_utils.AssertErrorsEqual(invalidPath, err, t)
}

func TestTransaction_GetByPathInvalidValue(t *testing.T) {
	_, err := GetByPath(10, "x")

	test_utils.AssertErrorsEqual(CannotAccessValueByPath, err, t)
}
