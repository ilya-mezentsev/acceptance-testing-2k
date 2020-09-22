package file_creator

import (
	"api_meta/models"
	"io/ioutil"
	"log"
	"os"
	"path"
	"services/errors"
	"services/plugins/response_factory"
	"services/tests_runner/plugins/tests_file_path"
	"test_utils"
	"testing"
	"utils"
)

var (
	s                       = New()
	testHash, testCasesPath string
	expectedSuccessStatus   = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus     = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	testCasesPath = path.Join(
		utils.MustGetEnv("TEST_CASES_ROOT_PATH"),
		testHash,
		utils.MustGetEnv("TEST_CASES_FILENAME"),
	)

	_ = os.RemoveAll(tests_file_path.BuildDirPath(testHash))
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_CreateTestsFileSuccess(t *testing.T) {
	response := s.CreateTestsFile(testHash, test_utils.MustGetFileUploadMockRequest(
		"tests_cases_file",
		testCasesPath,
		testHash,
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertTrue(
		test_utils.MustFileExists(tests_file_path.BuildFilePath(
			testHash,
			response.GetData().(models.CreateTestsFileResponse).Filename,
		)),
		t,
	)
}

func TestService_CreateTestsFileParseMultipartFormError(t *testing.T) {
	req := test_utils.MustGetFileUploadMockRequest(
		"tests_cases_file",
		testCasesPath,
		testHash,
	)
	req.Header.Del("Content-Type")

	response := s.CreateTestsFile(testHash, req)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestsFile,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		parseFormError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateTestsFileWrongFileKey(t *testing.T) {
	response := s.CreateTestsFile(testHash, test_utils.MustGetFileUploadMockRequest(
		"wrong_key",
		testCasesPath,
		testHash,
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestsFile,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		getFileFromRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateTestsFileCreateError(t *testing.T) {
	_ = os.Chmod(tests_file_path.BuildDirPath(testHash), 0)
	defer func() {
		_ = os.Chmod(tests_file_path.BuildDirPath(testHash), 0777)
	}()

	response := s.CreateTestsFile(testHash, test_utils.MustGetFileUploadMockRequest(
		"tests_cases_file",
		testCasesPath,
		testHash,
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestsFile,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		createTestFileError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateTmpFileError(t *testing.T) {
	_, err := s.createTmpFile(testHash, test_utils.BadReader())

	test_utils.AssertNotNil(err, t)
}
