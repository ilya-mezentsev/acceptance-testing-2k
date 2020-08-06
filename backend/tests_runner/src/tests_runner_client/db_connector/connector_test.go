package db_connector

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
	testDataPath, testHash string
)

func init() {
	testDataPath = path.Dir(test_utils.MustGetEnv("TEST_RUNNER_DB_FILE"))
	testHash = test_utils.MustGetEnv("TEST_ACCOUNT_HASH")
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestConnector_ConnectSuccess(t *testing.T) {
	connector := New(testDataPath)

	db, err := connector.Connect(testHash)

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(db, t)
	test_utils.AssertNil(db.Ping(), t)
}

func TestConnector_ConnectFileNotFound(t *testing.T) {
	connector := New("/home")

	_, err := connector.Connect(testHash)

	test_utils.AssertErrorsEqual(errors.DBFileNotFound, err, t)
}

func TestConnector_ConnectUnknownError(t *testing.T) {
	connector := New("/dev/null")

	_, err := connector.Connect(testHash)

	test_utils.AssertErrorsEqual(errors.UnknownError, err, t)
}
