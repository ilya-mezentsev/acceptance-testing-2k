package interfaces

import "models"

type TestsRunnerClient interface {
	Run(accountHash, testCasesFilename string) (models.TestsReport, models.ApplicationError)
}
