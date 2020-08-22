package tests_reader

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"test_utils"
	"testing"
	"utils"
)

var (
	testCasesRootPath, testCasesPath string
	testHash                         string
	expectedTestCasesData            string
)

func init() {
	testCasesRootPath = utils.MustGetEnv("TEST_CASES_ROOT_PATH")
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	testCasesPath = path.Join(
		testCasesRootPath,
		testHash,
		utils.MustGetEnv("TEST_CASES_FILENAME"),
	)

	data, err := ioutil.ReadFile(testCasesPath)
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
	data, err := Read(testCasesPath)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedTestCasesData, data, t)
}

func TestReader_ReadFileNotFound(t *testing.T) {
	_, err := Read(path.Join(testCasesRootPath, testHash, "not_exists.txt"))

	test_utils.AssertErrorsEqual(TestsFileNotFound, err, t)
}

func TestReader_ReadUnknownError(t *testing.T) {
	_, err := Read("/dev/null/tmp.txt")

	test_utils.AssertErrorsEqual(UnknownError, err, t)
}
