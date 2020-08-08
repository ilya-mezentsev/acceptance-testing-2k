package client

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http/httptest"
	"os"
	"path"
	mockCommand "test_runner_meta/mock/command"
	testRunnerClientMock "test_runner_meta/mock/test_runner_client"
	"test_utils"
	"testing"
	"tests_runner_client/errors"
	"utils"
)

var (
	testDataPath      string
	testCasesRootPath string
	testCasesFilename string
	testHash          string
	testCasesFilePath string
	db                *sqlx.DB
	r                 *mux.Router
	server            *httptest.Server
)

func init() {
	testDataPath = path.Dir(utils.MustGetEnv("TEST_RUNNER_DB_FILE"))
	testCasesRootPath = utils.MustGetEnv("TEST_CASES_ROOT_PATH")
	testCasesFilename = utils.MustGetEnv("TEST_CASES_FILENAME")
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")

	testCasesFilePath = path.Join(testCasesRootPath, testHash, testCasesFilename)
	r = mux.NewRouter()
	server = test_utils.GetTestServer(r)

	var err error
	db, err = sqlx.Open("sqlite3", path.Join(testDataPath, testHash, "db.db"))
	if err != nil {
		panic(err)
	}

	mockCommand.Init(r)
	testRunnerClientMock.FillTestCasesFile(testCasesFilePath)
	test_utils.ReplaceBaseURLAndInitTables(db, server.URL)
}

func TestMain(m *testing.M) {
	defer server.Close()
	defer test_utils.DropTables(db)

	os.Exit(m.Run())
}

func TestClient_RunSuccess(t *testing.T) {
	client := New(testDataPath, testCasesRootPath)

	report, err := client.Run(testHash, testCasesFilename)

	test_utils.AssertEqual(errors.EmptyApplicationError, err, t)
	test_utils.AssertEqual(testRunnerClientMock.PassedCount, report.PassedCount, t)
	test_utils.AssertEqual(testRunnerClientMock.FailedCount, report.FailedCount, t)
	test_utils.AssertEqual(testRunnerClientMock.FailedCount, len(report.Errors), t)
}

func TestClient_RunInvalidDBFilesRootPath(t *testing.T) {
	client := New("/home", testCasesRootPath)

	_, err := client.Run(testHash, testCasesFilename)

	test_utils.AssertEqual(unableToEstablishDBConnectionCode, err.Code, t)
}

func TestClient_RunInvalidTestCasesRootPath(t *testing.T) {
	client := New(testDataPath, "/home")

	_, err := client.Run(testHash, testCasesFilename)

	test_utils.AssertEqual(unableToReadTestCasesCode, err.Code, t)
}

func TestClient_RunNoDB(t *testing.T) {
	testRunnerClientMock.FillBadTestCasesData(testCasesFilePath)
	defer testRunnerClientMock.FillTestCasesFile(testCasesFilePath)

	client := New(testDataPath, testCasesRootPath)

	_, err := client.Run(testHash, testCasesFilename)

	test_utils.AssertEqual(unableToInitTestRunnersCode, err.Code, t)
}
