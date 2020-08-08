package services

import "api_meta/models"

func testObjectsToInterfaceSlice(testObjects []models.TestObject) []interface{} {
	var objects []interface{}
	for _, object := range testObjects {
		objects = append(objects, object)
	}

	return objects
}
