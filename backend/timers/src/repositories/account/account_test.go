package account

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"test_utils"
	"testing"
	"timers_meta/interfaces"
	"utils"
)

var (
	r  interfaces.AccountsRepository
	db *sqlx.DB
)

func init() {
	db = sqlx.MustOpen("sqlite3", utils.MustGetEnv("TEST_DB_FILE"))

	r = New(db)
}

func TestRepository_GetAccountsCreatedDayAgoSuccess(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)
	_ = db.MustExec(
		`INSERT INTO accounts(hash, created_at) VALUES(?, ?)`,
		"foo-bar-hash", "2020-09-30",
	)
	_ = db.MustExec(
		`INSERT INTO accounts(hash, created_at, verified) VALUES(?, ?, ?)`,
		"foo-bar-hash-2", "2020-09-29", true,
	)

	hashes, err := r.GetNonVerifiedAccountsCreatedDayAgo()

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(strings.Contains(
		strings.Join(hashes, "|"),
		"foo-bar-hash",
	), t)
	test_utils.AssertFalse(strings.Contains(
		strings.Join(hashes, "|"),
		"foo-bar-hash-2",
	), t)
}

func TestRepository_GetAccountsCreatedDayAgoNoTable(t *testing.T) {
	test_utils.DropTables(db)

	_, err := r.GetNonVerifiedAccountsCreatedDayAgo()

	test_utils.AssertNotNil(err, t)
}

func TestRepository_DeleteAccountsSuccess(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	err := r.DeleteAccounts([]string{test_utils.AccountHash})
	var accountExists bool
	_ = db.Get(&accountExists, `SELECT 1 FROM accounts WHERE hash = ?`, test_utils.AccountHash)

	test_utils.AssertNil(err, t)
	test_utils.AssertFalse(accountExists, t)
}

func TestRepository_DeleteAccountsNoTable(t *testing.T) {
	test_utils.DropTables(db)

	err := r.DeleteAccounts([]string{test_utils.AccountHash})

	test_utils.AssertNotNil(err, t)
}
