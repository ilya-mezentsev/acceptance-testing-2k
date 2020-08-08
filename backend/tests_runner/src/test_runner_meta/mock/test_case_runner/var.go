package test_case_runner

import "errors"

var (
	SomeTransactionError   = errors.New("some-error")
	SimpleMockTransaction  = MockTransaction{}
	ErroredMockTransaction = MockErroredTransaction{}
)
