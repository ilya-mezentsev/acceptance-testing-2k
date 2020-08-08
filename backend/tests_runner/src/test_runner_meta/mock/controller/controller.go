package controller

import (
	"test_runner_meta/models"
)

var (
	TestsReport = models.TestsReport{
		PassedCount: 2,
		FailedCount: 3,
		Errors: []models.TransactionError{
			{
				Code:            "some-code",
				Description:     "Some desc",
				TransactionText: "Some text",
			},
		},
	}
	ApplicationError = models.ApplicationError{
		Code:        "some-code",
		Description: "Some description",
	}
)

type calledWithContainer struct {
	accountHash, testCasesFilename string
}

func (c *calledWithContainer) Run(
	accountHash,
	testCasesFilename string,
) (models.TestsReport, models.ApplicationError) {
	c.accountHash = accountHash
	c.testCasesFilename = testCasesFilename

	return models.TestsReport{}, models.ApplicationError{}
}

func (c *calledWithContainer) CalledWith(accountHash, testCasesFilename string) bool {
	return c.accountHash == accountHash && c.testCasesFilename == testCasesFilename
}

type SimpleTestRunnerClientMock struct {
	calledWithContainer
}

type WithReportTestRunnerClientMock struct {
	calledWithContainer
}

func (m *WithReportTestRunnerClientMock) Run(
	accountHash,
	testCasesFilename string,
) (models.TestsReport, models.ApplicationError) {
	m.calledWithContainer.Run(accountHash, testCasesFilename)

	return TestsReport, models.ApplicationError{}
}

type WithApplicationErrorTestRunnerClientMock struct {
	calledWithContainer
}

func (m *WithApplicationErrorTestRunnerClientMock) Run(
	accountHash,
	testCasesFilename string,
) (models.TestsReport, models.ApplicationError) {
	m.calledWithContainer.Run(accountHash, testCasesFilename)

	return models.TestsReport{}, ApplicationError
}
