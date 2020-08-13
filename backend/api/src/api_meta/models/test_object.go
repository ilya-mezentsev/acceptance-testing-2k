package models

type (
	TestObject struct {
		Name string `json:"name" validation:"regular-name"`
		Hash string `json:"hash" validation:"md5-hash"`
	}

	CreateTestObjectRequest struct {
		AccountHash string     `json:"account_hash" validation:"md5-hash"`
		TestObject  TestObject `json:"test_object"`
	}
)
