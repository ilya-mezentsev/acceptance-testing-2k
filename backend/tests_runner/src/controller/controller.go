package controller

import (
	"context"
	"interfaces"
	"models"
	"test_case_runner"
	"tests_runner_client/errors"
)

type Controller struct {
	testsRunnerClient interfaces.TestsRunnerClient
}

func New(testsRunnerClient interfaces.TestsRunnerClient) *Controller {
	return &Controller{testsRunnerClient}
}

func (c *Controller) Run(
	_ context.Context,
	request *test_case_runner.TestCasesRequest,
) (*test_case_runner.TestsReport, error) {
	report, applicationError := c.testsRunnerClient.Run(request.AccountHash, request.TestCasesFilename)
	if applicationError != errors.EmptyApplicationError {
		return &test_case_runner.TestsReport{
			ApplicationError: &test_case_runner.ApplicationError{
				Code:        applicationError.Code,
				Description: applicationError.Description,
			},
		}, nil
	}

	return &test_case_runner.TestsReport{
		Report: &test_case_runner.TestCaseRunReport{
			PassedCount: int64(report.PassedCount),
			FailedCount: int64(report.FailedCount),
			Errors:      castReportErrors(report.Errors),
		},
	}, nil
}

func castReportErrors(errors []models.TransactionError) []*test_case_runner.TransactionError {
	var castedErrors []*test_case_runner.TransactionError
	for _, err := range errors {
		castedErrors = append(castedErrors, &test_case_runner.TransactionError{
			Code:            err.Code,
			Description:     err.Description,
			TransactionText: err.TransactionText,
		})
	}

	return castedErrors
}
