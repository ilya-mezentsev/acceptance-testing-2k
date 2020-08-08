package interfaces

import "test_runner_meta/models"

type TestsRunnerClient interface {
	Run(accountHash, testCasesFilename string) (models.TestsReport, models.ApplicationError)
}
