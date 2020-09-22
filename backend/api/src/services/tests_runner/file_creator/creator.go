package file_creator

import (
	"api_meta/interfaces"
	"api_meta/models"
	"io"
	"logger"
	"net/http"
	"os"
	"services/errors"
	"services/plugins/hash"
	"services/plugins/response_factory"
	"services/tests_runner/plugins/tests_file_path"
)

type Service struct {
}

func New() Service {
	return Service{}
}

func (s Service) CreateTestsFile(accountHash string, r *http.Request) interfaces.Response {
	err := r.ParseMultipartForm(defaultMaxMemory)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestsFile,
			Description: parseFormError,
		})
	}

	file, _, err := r.FormFile("tests_cases_file")
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestsFile,
			Description: getFileFromRequestError,
		})
	}
	defer func() {
		_ = file.Close()
	}()

	testCasesFilename, err := s.createTmpFile(accountHash, file)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestsFile,
			Description: createTestFileError,
		})
	}

	return response_factory.SuccessResponse(models.CreateTestsFileResponse{
		Filename: testCasesFilename,
	})
}

func (s Service) createTmpFile(
	accountHash string,
	fileFromRequest io.Reader,
) (string, error) {
	accountTmpFilesRootPath := tests_file_path.BuildDirPath(accountHash)
	_, err := os.Stat(accountTmpFilesRootPath)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(accountTmpFilesRootPath, 0777)
	}

	testCasesFilename := "test_cases" + hash.Md5WithTimeAsKey("test_cases") + ".txt"
	testCasesPath := tests_file_path.BuildFilePath(accountHash, testCasesFilename)
	f, err := os.OpenFile(testCasesPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to open file: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"file_path": testCasesPath,
			},
		}, logger.Error)

		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = io.Copy(f, fileFromRequest)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to copy to test cases file: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"file_path": testCasesPath,
			},
		}, logger.Error)

		return "", err
	}

	return testCasesFilename, nil
}
