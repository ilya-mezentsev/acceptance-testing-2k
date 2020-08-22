package runner

import (
	"api_meta/mock/services"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"services/errors"
	"services/plugins/response_factory"
	"services/tests_runner/client"
	"services/tests_runner/tests_file"
	"test_case_runner"
	"test_utils"
	"testing"
	"time"
	"utils"
)

var (
	filesRootPath, testHash string
	testCasesPath           string
	filesManager            tests_file.Manager
	serviceServerMock       = services.GRPCServiceServerMock{}
	expectedSuccessStatus   = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus     = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	filesRootPath = utils.MustGetEnv("TEST_CASES_ROOT_PATH")
	testCasesPath = path.Join(
		filesRootPath,
		testHash,
		utils.MustGetEnv("TEST_CASES_FILENAME"),
	)
	filesManager = tests_file.New(filesRootPath)
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	res := m.Run()
	files, _ := filepath.Glob(path.Join(filesRootPath, testHash, "test_cases_*.txt"))
	for _, file := range files {
		_ = os.Remove(file)
	}
	os.Exit(res)
}

func TestService_RunSuccess(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	c := client.New(serviceAddress)
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	response := New(filesManager, c).Run(
		testHash,
		test_utils.MustGetFileUploadMockRequest("tests_cases_file", testCasesPath),
	)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		services.MockTestCasesReport.Report.PassedCount,
		response.GetData().(*test_case_runner.TestsReport).Report.PassedCount,
		t,
	)
	test_utils.AssertEqual(
		services.MockTestCasesReport.Report.FailedCount,
		response.GetData().(*test_case_runner.TestsReport).Report.FailedCount,
		t,
	)
}

func TestService_RunParseMultipartFormError(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	c := client.New(serviceAddress)
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	req := test_utils.MustGetFileUploadMockRequest("tests_cases_file", testCasesPath)
	req.Header.Del("Content-Type")
	response := New(filesManager, c).Run(testHash, req)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToRunTestsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		parseFormError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_RunWrongFileKey(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	c := client.New(serviceAddress)
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	response := New(filesManager, c).Run(
		testHash,
		test_utils.MustGetFileUploadMockRequest("wrong_key", testCasesPath),
	)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToRunTestsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		getFileFromRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_RunCreateFileError(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	c := client.New(serviceAddress)
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	response := New(tests_file.New("/"), c).Run(
		testHash,
		test_utils.MustGetFileUploadMockRequest("tests_cases_file", testCasesPath),
	)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToRunTestsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		createTestFileError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_RunCalGRPCServiceError(t *testing.T) {
	c := client.New("")
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	response := New(filesManager, c).Run(
		testHash,
		test_utils.MustGetFileUploadMockRequest("tests_cases_file", testCasesPath),
	)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToRunTestsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		callRemoteProcedureError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_SilentlyRemoveNotExistsFile(t *testing.T) {
	New(filesManager, client.New("")).silentlyRemoveTestCasesFile("/blah-blah")
}
