package tests_file

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"test_utils"
	"testing"
	"utils"
)

var (
	filesRootPath, testHash string
	filesManager            Manager
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	filesRootPath = utils.MustGetEnv("TEST_CASES_ROOT_PATH")
	filesManager = New(filesRootPath)
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

func TestManager_CreateTestCasesFileSuccess(t *testing.T) {
	filename, err := filesManager.CreateTestCasesFile(testHash, test_utils.GetReadCloser("Hello"))

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(test_utils.MustFileExists(filename), t)

	data, _ := ioutil.ReadFile(filename)
	test_utils.AssertEqual("Hello", string(data), t)
}

func TestManager_CreateTestCasesFileErrorPermissions(t *testing.T) {
	filesManager := New("/")

	_, err := filesManager.CreateTestCasesFile(testHash, test_utils.GetReadCloser("Hello"))
	test_utils.AssertNotNil(err, t)
}

func TestManager_CreateTestCasesFileErrorBadReader(t *testing.T) {
	_, err := filesManager.CreateTestCasesFile(testHash, test_utils.BadReader())

	test_utils.AssertNotNil(err, t)
}

func TestManager_RemoveFile(t *testing.T) {
	filename, _ := filesManager.CreateTestCasesFile(testHash, test_utils.GetReadCloser("Hello"))

	test_utils.AssertTrue(test_utils.MustFileExists(filename), t)

	err := filesManager.RemoveFile(filename)

	test_utils.AssertNil(err, t)
	test_utils.AssertFalse(test_utils.MustFileExists(filename), t)
}
