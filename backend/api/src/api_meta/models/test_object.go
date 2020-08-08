package models

type (
	TestObject struct {
		Name string `json:"name"`
		Hash string `json:"hash"`
	}

	CreateTestObjectRequest struct {
		AccountHash string     `json:"account_hash"`
		TestObject  TestObject `json:"test_object"`
	}

	UpdateTestObjectRequest struct {
		AccountHash   string        `json:"account_hash"`
		UpdatePayload []UpdateModel `json:"update_payload"`
	}
)
