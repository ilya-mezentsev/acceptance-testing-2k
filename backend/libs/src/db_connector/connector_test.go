package db_connector

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"test_utils"
	"testing"
	"time"
	"utils"
)

var (
	testDataPath, testHash string
)

func init() {
	testDataPath = path.Dir(utils.MustGetEnv("TEST_DB_FILE"))
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
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

	test_utils.AssertErrorsEqual(DBFileNotFound, err, t)
}

func TestConnector_ConnectUnknownError(t *testing.T) {
	connector := New("/dev/null")

	_, err := connector.Connect(testHash)

	test_utils.AssertErrorsEqual(UnknownError, err, t)
}

func TestConnector_ConnectExpirationCheck(t *testing.T) {
	connectionCacheCleanTimout = time.Millisecond * 20
	connectionCacheLifetime = time.Microsecond

	connector := New(testDataPath)
	db, err := connector.Connect(testHash)

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(db, t)
	test_utils.AssertNil(db.Ping(), t)

	time.Sleep(connectionCacheCleanTimout * 2)

	_, found := connector.accountHashToConnection[testHash]

	test_utils.AssertFalse(found, t)
}
