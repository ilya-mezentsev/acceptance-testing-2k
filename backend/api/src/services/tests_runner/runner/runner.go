package runner

import (
	"api_meta/interfaces"
	"logger"
	"net/http"
	"services/errors"
	"services/plugins/response_factory"
	"services/tests_runner/client"
	"services/tests_runner/tests_file"
)

type Service struct {
	client           client.Grpc
	testsFileManager tests_file.Manager
}

func New(
	testsFileManager tests_file.Manager,
	client client.Grpc,
) Service {
	return Service{
		testsFileManager: testsFileManager,
		client:           client,
	}
}

func (s Service) Run(accountHash string, r *http.Request) interfaces.Response {
	err := r.ParseMultipartForm(defaultMaxMemory)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToRunTestsCode,
			Description: parseFormError,
		})
	}

	file, _, err := r.FormFile("tests_cases_file")
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToRunTestsCode,
			Description: getFileFromRequestError,
		})
	}
	defer func() {
		_ = file.Close()
	}()

	testCasesFilePath, err := s.testsFileManager.CreateTestCasesFile(accountHash, file)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToRunTestsCode,
			Description: createTestFileError,
		})
	}
	defer s.silentlyRemoveTestCasesFile(testCasesFilePath)

	testsReport, err := s.client.Call(accountHash, testCasesFilePath)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToRunTestsCode,
			Description: callRemoteProcedureError,
		})
	}

	return response_factory.SuccessResponse(testsReport)
}

func (s Service) silentlyRemoveTestCasesFile(filePath string) {
	err := s.testsFileManager.RemoveFile(filePath)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to remove file: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"file_path": filePath,
			},
		}, logger.Error)
	}
}
