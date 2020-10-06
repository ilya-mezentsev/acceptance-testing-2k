package registration

import (
	"api_meta/interfaces"
	"api_meta/models"
	"api_meta/types"
	"db_connector"
	"errors"
	"github.com/jmoiron/sqlx"
	"path"
	"services/plugins/hash"
	"test_utils"
	"testing"
	"utils"
)

var (
	testHash  string
	db        *sqlx.DB
	connector *db_connector.Connector
	r         interfaces.RegistrationRepository
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)

	r = New(db, connector)
}

func TestRepository_CreateAccountSuccess(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	accountHash := hash.Md5WithTimeAsKey("some-hash")
	err := r.CreateAccount(accountHash)

	var accountCreated bool
	_ = db.Get(&accountCreated, `SELECT 1 FROM accounts WHERE hash = ?`, accountHash)
	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(accountCreated, t)
}

func TestRepository_CreateAccountAlreadyExistsError(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	err := r.CreateAccount(test_utils.AccountHash)

	var expectedError types.AccountHashAlreadyExists
	test_utils.AssertTrue(errors.As(err, &expectedError), t)
}

func TestRepository_CreateAccountNoTableError(t *testing.T) {
	test_utils.DropTables(db)

	err := r.CreateAccount("foo")

	var notExpectedError types.AccountHashAlreadyExists
	test_utils.AssertNotNil(err, t)
	test_utils.AssertFalse(errors.As(err, &notExpectedError), t)
}

func TestRepository_CreateAccountCredentialsSuccess(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	err := r.CreateAccountCredentials(models.AccountCredentialsRecord{
		AccountHash: testHash,
		Login:       "some-login",
		Password:    "some-password",
	})

	var credentialsCreated bool
	_ = db.Get(
		&credentialsCreated,
		`SELECT 1 FROM account_credentials
		WHERE login = 'some-login' AND password = 'some-password'`,
		testHash,
	)
	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(credentialsCreated, t)
}

func TestRepository_CreateAccountCredentialsBadAccountHashError(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	err := r.CreateAccountCredentials(models.AccountCredentialsRecord{
		AccountHash: hash.Md5WithTimeAsKey("bad-hash"),
		Login:       "some-login",
		Password:    "some-password",
	})

	test_utils.AssertNotNil(err, t)
}
