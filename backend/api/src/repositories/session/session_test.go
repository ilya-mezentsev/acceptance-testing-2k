package session

import (
	"api_meta/interfaces"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"test_utils"
	"testing"
	"utils"
)

var (
	db *sqlx.DB
	r  interfaces.SessionRepository
)

func init() {
	dbFile := utils.MustGetEnv("TEST_DB_FILE")

	var err error
	db, err = sqlx.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	r = New(db)
}

func TestRepository_AccountExistsTrue(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	accountExists, err := r.AccountExists(test_utils.AccountHash)

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(accountExists, t)
}

func TestRepository_AccountExistsFalse(t *testing.T) {
	test_utils.InitTablesWithAccounts(db)
	defer test_utils.DropTables(db)

	accountExists, err := r.AccountExists("blah-blah")

	test_utils.AssertNil(err, t)
	test_utils.AssertFalse(accountExists, t)
}

func TestRepository_AccountExistsError(t *testing.T) {
	test_utils.DropTables(db)

	_, err := r.AccountExists(test_utils.AccountHash)

	test_utils.AssertNotNil(err, t)
}
