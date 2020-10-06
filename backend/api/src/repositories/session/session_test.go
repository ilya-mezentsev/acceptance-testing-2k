package session

import (
	"api_meta/interfaces"
	"db_connector"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"test_utils"
	"testing"
	"utils"
)

var (
	testHash  string
	db        *sqlx.DB
	connector *db_connector.Connector
	r         interfaces.SessionRepository
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)

	r = New(connector)
}

func TestRepository_CredentialsExistsTrue(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	credentialsExists, err := r.CredentialsExists(
		testHash,
		test_utils.CredentialsLogin,
		test_utils.CredentialsPassword,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(credentialsExists, t)
}

func TestRepository_CredentialsExistsFalse(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	credentialsExists, err := r.CredentialsExists(
		testHash,
		"foo",
		"bar",
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertFalse(credentialsExists, t)
}

func TestRepository_CredentialsExistsBadAccountHash(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	_, err := r.CredentialsExists(
		"blah-blah",
		test_utils.CredentialsLogin,
		test_utils.CredentialsPassword,
	)

	test_utils.AssertErrorsEqual(db_connector.DBFileNotFound, err, t)
}

func TestRepository_CredentialsExistsError(t *testing.T) {
	test_utils.DropTables(db)

	_, err := r.CredentialsExists(
		testHash,
		test_utils.CredentialsLogin,
		test_utils.CredentialsPassword,
	)

	test_utils.AssertNotNil(err, t)
}
