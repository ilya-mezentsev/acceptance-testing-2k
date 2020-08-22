package services

import (
	"context"
	"test_case_runner"
)

type GRPCServiceServerMock struct {
}

func (m GRPCServiceServerMock) Run(
	_ context.Context,
	request *test_case_runner.TestCasesRequest,
) (*test_case_runner.TestsReport, error) {
	if request.AccountHash == BadAccountHash {
		return nil, someError
	}

	return MockTestCasesReport, nil
}
