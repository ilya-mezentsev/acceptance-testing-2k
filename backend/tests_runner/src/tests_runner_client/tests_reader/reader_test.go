package tests_reader

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"test_utils"
	"testing"
	"tests_runner_client/errors"
)

var (
	testCasesRootPath, testCasesFilename, testHash string
	expectedTestCasesData                          string
)

func init() {
	testCasesRootPath = test_utils.MustGetEnv("TEST_CASES_ROOT_PATH")
	testCasesFilename = test_utils.MustGetEnv("TEST_CASES_FILENAME")
	testHash = test_utils.MustGetEnv("TEST_ACCOUNT_HASH")

	data, err := ioutil.ReadFile(path.Join(testCasesRootPath, testHash, testCasesFilename))
	if err != nil {
		panic(err)
	}
	expectedTestCasesData = string(data)
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestReader_ReadSuccess(t *testing.T) {
	reader := New(testCasesRootPath)

	data, err := reader.Read(testHash, testCasesFilename)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedTestCasesData, data, t)
}

func TestReader_ReadFileNotFound(t *testing.T) {
	reader := New(testCasesRootPath)

	_, err := reader.Read(testHash, "not_exists.txt")

	test_utils.AssertErrorsEqual(errors.TestsFileNotFound, err, t)
}

func TestReader_ReadUnknownError(t *testing.T) {
	reader := New("/dev/null")

	_, err := reader.Read(testHash, testCasesFilename)

	test_utils.AssertErrorsEqual(errors.UnknownError, err, t)
}
