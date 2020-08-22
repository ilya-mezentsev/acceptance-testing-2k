package tests_runner

import (
	"api_meta/mock/services"
	"controllers/plugins/response_writer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"services/plugins/hash"
	"services/plugins/response_factory"
	"services/tests_runner/client"
	"services/tests_runner/runner"
	"services/tests_runner/tests_file"
	"test_utils"
	"testing"
	"time"
	"utils"
)

var (
	r                       *mux.Router
	filesRootPath, testHash string
	testCasesPath           string
	filesManager            tests_file.Manager
	serviceServerMock       = services.GRPCServiceServerMock{}
	expectedSuccessStatus   = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus     = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	r = mux.NewRouter()
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

func TestRun_Success(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	runnerService := runner.New(filesManager, client.New(serviceAddress))
	Init(r, runnerService)

	server := test_utils.GetTestServer(r)
	defer server.Close()

	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	req := test_utils.MustGetFileUploadRequest(
		fmt.Sprintf("%s/tests/%s", server.URL, testHash),
		"tests_cases_file",
		testCasesPath,
	)
	responseData := test_utils.MustDoRequest(req)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedSuccessStatus, response.Status, t)
}

func TestRun_Error(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	runnerService := runner.New(filesManager, client.New(serviceAddress))
	Init(r, runnerService)

	server := test_utils.GetTestServer(r)
	defer server.Close()

	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	req := test_utils.MustGetFileUploadRequest(
		fmt.Sprintf("%s/tests/%s", server.URL, hash.Md5(testHash)),
		"tests_cases_file",
		testCasesPath,
	)
	responseData := test_utils.MustDoRequest(req)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedErrorStatus, response.Status, t)
}
