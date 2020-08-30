package test_command_meta

import (
	"api_meta/models"
	"db_connector"
	"github.com/jmoiron/sqlx"
	"path"
	"test_utils"
	"testing"
	"utils"
)

var (
	testHash  string
	db        *sqlx.DB
	connector db_connector.Connector
	r         Repository
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)

	r = New(connector)
}

func TestRepository_CreateSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.Create(testHash, models.CommandKeyValue{
		Headers: []models.KeyValueMapping{
			{
				Hash:        "hash-1",
				Key:         "key-1",
				Value:       "value-1",
				CommandHash: test_utils.ObjectHash,
			},
		},
		Cookies: []models.KeyValueMapping{
			{
				Hash:        "hash-2",
				Key:         "key-2",
				Value:       "value-2",
				CommandHash: test_utils.ObjectHash,
			},
		},
	})

	var (
		createdHeader models.KeyValueMapping
		createdCookie models.KeyValueMapping
	)

	_ = db.Get(
		&createdHeader,
		`SELECT key, value FROM commands_headers WHERE hash = $1 AND command_hash = $2`,
		"hash-1", test_utils.ObjectHash,
	)
	_ = db.Get(
		&createdCookie,
		`SELECT key, value FROM commands_cookies WHERE hash = $1 AND command_hash = $2`,
		"hash-2", test_utils.ObjectHash,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("key-1", createdHeader.Key, t)
	test_utils.AssertEqual("value-1", createdHeader.Value, t)
	test_utils.AssertEqual("key-2", createdCookie.Key, t)
	test_utils.AssertEqual("value-2", createdCookie.Value, t)
}

func TestRepository_CreateBadAccountHash(t *testing.T) {
	err := r.Create("bad-hash", models.CommandKeyValue{})

	test_utils.AssertNotNil(err, t)
}

func TestRepository_CreateNoHeadersTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	err := r.Create(testHash, models.CommandKeyValue{
		Headers: []models.KeyValueMapping{
			{
				Hash:        "hash-1",
				Key:         "key-1",
				Value:       "value-1",
				CommandHash: test_utils.ObjectHash,
			},
		},
	})

	test_utils.AssertNotNil(err, t)
}

func TestRepository_CreateNoCookiesTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsCookies(db)
	defer test_utils.DropTables(db)

	err := r.Create(testHash, models.CommandKeyValue{
		Cookies: []models.KeyValueMapping{
			{
				Hash:        "hash-2",
				Key:         "key-2",
				Value:       "value-2",
				CommandHash: test_utils.ObjectHash,
			},
		},
	})

	test_utils.AssertNotNil(err, t)
}

func TestRepository_UpdateHeadersSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.UpdateHeaders(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.HeaderHash1,
			FieldName: "key",
			NewValue:  "FOO",
		},
	})

	var updatedHeader models.KeyValueMapping
	_ = db.Get(
		&updatedHeader,
		`SELECT key, value FROM commands_headers WHERE hash = $1 AND command_hash = $2`,
		test_utils.HeaderHash1, test_utils.ObjectHash,
	)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("FOO", updatedHeader.Key, t)
}

func TestRepository_UpdateHeadersBadHash(t *testing.T) {
	err := r.UpdateHeaders("bad-hash", nil)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_UpdateHeadersNoHeadersTables(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	err := r.UpdateHeaders(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.HeaderHash1,
			FieldName: "key",
			NewValue:  "FOO",
		},
	})

	test_utils.AssertNotNil(err, t)
}

func TestRepository_UpdateCookiesSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.UpdateCookies(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.CookieHash1,
			FieldName: "key",
			NewValue:  "FOO",
		},
	})

	var updatedCookie models.KeyValueMapping
	_ = db.Get(
		&updatedCookie,
		`SELECT key, value FROM commands_cookies WHERE hash = $1 AND command_hash = $2`,
		test_utils.CookieHash1, test_utils.ObjectHash,
	)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("FOO", updatedCookie.Key, t)
}

func TestRepository_DeleteHeaderSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.DeleteHeader(testHash, test_utils.HeaderHash1)

	var headerExists bool
	_ = db.Get(
		&headerExists,
		`SELECT 1 FROM commands_headers WHERE hash = ?`,
		test_utils.HeaderHash1,
	)
	test_utils.AssertNil(err, t)
	test_utils.AssertFalse(headerExists, t)
}

func TestRepository_DeleteHeaderBadAccountHash(t *testing.T) {
	err := r.DeleteHeader("bad-hash", test_utils.HeaderHash1)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_DeleteCookieSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.DeleteCookie(testHash, test_utils.CookieHash1)

	var cookieExists bool
	_ = db.Get(
		&cookieExists,
		`SELECT 1 FROM commands_cookies WHERE hash = ?`,
		test_utils.CookieHash1,
	)
	test_utils.AssertNil(err, t)
	test_utils.AssertFalse(cookieExists, t)
}
