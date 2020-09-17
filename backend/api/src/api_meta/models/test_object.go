package models

type (
	TestObject struct {
		Name string `json:"name" validation:"regular-name"`
		Hash string `json:"hash" validation:"md5-hash"`
	}

	CreateTestObjectRequest struct {
		TestObject TestObject `json:"test_object"`
	}
)
