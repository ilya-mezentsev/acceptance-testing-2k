package client

import (
	"command/http/builder"
	"db_connector"
	"github.com/jmoiron/sqlx"
	"test_case/factory"
	"test_runner_meta/interfaces"
	"test_runner_meta/models"
	"tests_runner_client/errors"
	"tests_runner_client/tests_reader"
)

const (
	unableToInitTestRunnersCode       = "unable-to-init-tests-runners"
	unableToEstablishDBConnectionCode = "unable-establish-db-connection"
	unableToReadTestCasesCode         = "unable-to-read-test-cases"
)

type Client struct {
	dbConnector *db_connector.Connector
}

func New(dbRootPath string) interfaces.TestsRunnerClient {
	return Client{db_connector.New(dbRootPath)}
}

func (c Client) Run(
	accountHash,
	testCasesPath string,
) (models.TestsReport, models.ApplicationError) {
	db, testCases, applicationError := c.getDBConnectionAndTestCasesData(
		accountHash,
		testCasesPath,
	)
	if applicationError != errors.EmptyApplicationError {
		return models.TestsReport{}, applicationError
	}

	testsRunners, err := factory.New(builder.New(db)).GetAll(testCases)
	if err != nil {
		return models.TestsReport{}, models.ApplicationError{
			Code:        unableToInitTestRunnersCode,
			Description: err.Error(),
		}
	}

	return c.runTests(testsRunners), errors.EmptyApplicationError
}

func (c Client) getDBConnectionAndTestCasesData(
	accountHash,
	testCasesPath string,
) (*sqlx.DB, string, models.ApplicationError) {
	db, err := c.dbConnector.Connect(accountHash)
	if err != nil {
		return nil, "", models.ApplicationError{
			Code:        unableToEstablishDBConnectionCode,
			Description: err.Error(),
		}
	}

	testCases, err := tests_reader.Read(testCasesPath)
	if err != nil {
		return nil, "", models.ApplicationError{
			Code:        unableToReadTestCasesCode,
			Description: err.Error(),
		}
	}

	return db, testCases, errors.EmptyApplicationError
}

func (c Client) runTests(testsRunners []interfaces.TestCaseRunner) models.TestsReport {
	var testsReport models.TestsReport
	testsCount := len(testsRunners)
	processing := models.TestsRun{
		Success: make(chan bool),
		Error:   make(chan models.TransactionError),
	}

	for _, testsRunner := range testsRunners {
		go testsRunner.Run(processing)
	}

	for {
		select {
		case err := <-processing.Error:
			testsReport.FailedCount++
			testsReport.Errors = append(testsReport.Errors, err)
		case <-processing.Success:
			testsReport.PassedCount++
		default:
			if (testsReport.PassedCount + testsReport.FailedCount) >= testsCount {
				return testsReport
			}
		}
	}
}
