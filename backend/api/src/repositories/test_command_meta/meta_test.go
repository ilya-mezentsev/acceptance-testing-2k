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
	connector *db_connector.Connector
	r         Repository
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)

	r = New(connector)
}

func TestRepository_GetAllHeadersAndCookiesSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	headers, cookies, err := r.GetAllHeadersAndCookies(testHash)

	test_utils.AssertNil(err, t)
	for i, expectedHeader := range test_utils.Headers {
		test_utils.AssertEqual(expectedHeader["key"], headers[i].Key, t)
		test_utils.AssertEqual(expectedHeader["value"], headers[i].Value, t)
		test_utils.AssertEqual(expectedHeader["hash"], headers[i].Hash, t)
		test_utils.AssertEqual(expectedHeader["command_hash"], headers[i].CommandHash, t)
	}

	for i, expectedCookie := range test_utils.Cookies {
		test_utils.AssertEqual(expectedCookie["key"], cookies[i].Key, t)
		test_utils.AssertEqual(expectedCookie["value"], cookies[i].Value, t)
		test_utils.AssertEqual(expectedCookie["hash"], cookies[i].Hash, t)
		test_utils.AssertEqual(expectedCookie["command_hash"], cookies[i].CommandHash, t)
	}
}

func TestRepository_GetAllHeadersAndCookiesBadAccountHash(t *testing.T) {
	_, _, err := r.GetAllHeadersAndCookies("bad-hash")

	test_utils.AssertNotNil(err, t)
}

func TestRepository_GetAllHeadersAndCookiesNoHeadersTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	_, _, err := r.GetAllHeadersAndCookies(testHash)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_GetAllHeadersAndCookiesNoCookiesTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsCookies(db)
	defer test_utils.DropTables(db)

	_, _, err := r.GetAllHeadersAndCookies(testHash)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_GetCommandHeadersAndCookiesSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	headers, cookies, err := r.GetCommandHeadersAndCookies(
		testHash,
		test_utils.CreateCommandHash,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.Headers[0]["key"], headers[0].Key, t)
	test_utils.AssertEqual(test_utils.Headers[0]["value"], headers[0].Value, t)
	test_utils.AssertEqual(test_utils.Headers[0]["hash"], headers[0].Hash, t)
	test_utils.AssertEqual(test_utils.Headers[0]["command_hash"], headers[0].CommandHash, t)

	test_utils.AssertEqual(test_utils.Cookies[0]["key"], cookies[0].Key, t)
	test_utils.AssertEqual(test_utils.Cookies[0]["value"], cookies[0].Value, t)
	test_utils.AssertEqual(test_utils.Cookies[0]["hash"], cookies[0].Hash, t)
	test_utils.AssertEqual(test_utils.Cookies[0]["command_hash"], cookies[0].CommandHash, t)
}

func TestRepository_GetCommandHeadersAndCookiesBadAccountHash(t *testing.T) {
	_, _, err := r.GetCommandHeadersAndCookies(
		"bad-hash",
		test_utils.CreateCommandHash,
	)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_GetCommandHeadersAndCookiesNoHeadersTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	_, _, err := r.GetCommandHeadersAndCookies(
		testHash,
		test_utils.CreateCommandHash,
	)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_GetCommandHeadersAndCookiesNoCookiesTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsCookies(db)
	defer test_utils.DropTables(db)

	_, _, err := r.GetCommandHeadersAndCookies(
		testHash,
		test_utils.CreateCommandHash,
	)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_CreateSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.Create(testHash, models.CommandMeta{
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

func TestRepository_CreateFewMetaSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.Create(testHash, models.CommandMeta{
		Headers: []models.KeyValueMapping{
			{
				Hash:        "hash-1",
				Key:         "key-1",
				Value:       "value-1",
				CommandHash: test_utils.ObjectHash,
			},
			{
				Hash:        "hash-2",
				Key:         "key-2",
				Value:       "value-2",
				CommandHash: test_utils.ObjectHash,
			},
			{
				Hash:        test_utils.HeaderHash1,
				Key:         test_utils.HeaderKey1,
				Value:       "FOO_BAR",
				CommandHash: test_utils.ObjectHash,
			},
		},
		Cookies: []models.KeyValueMapping{
			{
				Hash:        "hash-3",
				Key:         "key-3",
				Value:       "value-3",
				CommandHash: test_utils.ObjectHash,
			},
			{
				Hash:        "hash-4",
				Key:         "key-4",
				Value:       "value-4",
				CommandHash: test_utils.ObjectHash,
			},
			{
				Hash:        test_utils.CookieHash1,
				Key:         test_utils.CookieKey1,
				Value:       "FOO_BAR",
				CommandHash: test_utils.ObjectHash,
			},
		},
	})

	var (
		header1Created, header2Created, existsHeaderReplaced bool
		cookie1Created, cookie2Created, existsCookieReplaced bool
	)
	_ = db.Get(
		&header1Created,
		`SELECT 1 FROM commands_headers WHERE hash = $1 AND command_hash = $2`,
		"hash-1", test_utils.ObjectHash,
	)
	_ = db.Get(
		&header2Created,
		`SELECT 1 FROM commands_headers WHERE hash = $1 AND command_hash = $2`,
		"hash-2", test_utils.ObjectHash,
	)
	_ = db.Get(
		&existsHeaderReplaced,
		`SELECT 1 FROM commands_headers WHERE hash = $1 AND command_hash = $2`,
		test_utils.HeaderHash1, test_utils.ObjectHash,
	)
	_ = db.Get(
		&cookie1Created,
		`SELECT 1 FROM commands_cookies WHERE hash = $1 AND command_hash = $2`,
		"hash-3", test_utils.ObjectHash,
	)
	_ = db.Get(
		&cookie2Created,
		`SELECT 1 FROM commands_cookies WHERE hash = $1 AND command_hash = $2`,
		"hash-4", test_utils.ObjectHash,
	)
	_ = db.Get(
		&existsCookieReplaced,
		`SELECT 1 FROM commands_cookies WHERE hash = $1 AND command_hash = $2`,
		test_utils.CookieHash1, test_utils.ObjectHash,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(header1Created, t)
	test_utils.AssertTrue(header2Created, t)
	test_utils.AssertTrue(existsHeaderReplaced, t)
	test_utils.AssertTrue(cookie1Created, t)
	test_utils.AssertTrue(cookie2Created, t)
	test_utils.AssertTrue(existsCookieReplaced, t)
}

func TestRepository_CreateBadAccountHash(t *testing.T) {
	err := r.Create("bad-hash", models.CommandMeta{})

	test_utils.AssertNotNil(err, t)
}

func TestRepository_CreateNoHeadersTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	err := r.Create(testHash, models.CommandMeta{
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

	err := r.Create(testHash, models.CommandMeta{
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

func TestRepository_UpdateHeadersAndCookiesSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.UpdateHeadersAndCookies(
		testHash,
		[]models.UpdateModel{
			{
				Hash:      test_utils.HeaderHash1,
				FieldName: "key",
				NewValue:  "FOO",
			},
		},
		[]models.UpdateModel{
			{
				Hash:      test_utils.CookieHash1,
				FieldName: "value",
				NewValue:  "BAR",
			},
		},
	)

	var headerUpdated bool
	_ = db.Get(
		&headerUpdated,
		`SELECT 1 FROM commands_headers WHERE key = $1 AND hash = $2`,
		"FOO", test_utils.HeaderHash1,
	)

	var cookieUpdated bool
	_ = db.Get(
		&cookieUpdated,
		`SELECT 1 FROM commands_cookies WHERE value = $1 AND hash = $2`,
		"BAR", test_utils.CookieHash1,
	)

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(headerUpdated, t)
	test_utils.AssertTrue(cookieUpdated, t)
}

func TestRepository_UpdateHeadersAndCookiesBadAccountHash(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.UpdateHeadersAndCookies(
		"bad-hash",
		[]models.UpdateModel{
			{
				Hash:      test_utils.HeaderHash1,
				FieldName: "key",
				NewValue:  "FOO",
			},
		},
		[]models.UpdateModel{
			{
				Hash:      test_utils.CookieHash1,
				FieldName: "value",
				NewValue:  "BAR",
			},
		},
	)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_UpdateHeadersAndCookiesNoHeadersTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	err := r.UpdateHeadersAndCookies(
		"bad-hash",
		[]models.UpdateModel{
			{
				Hash:      test_utils.HeaderHash1,
				FieldName: "key",
				NewValue:  "FOO",
			},
		},
		[]models.UpdateModel{
			{
				Hash:      test_utils.CookieHash1,
				FieldName: "value",
				NewValue:  "BAR",
			},
		},
	)

	test_utils.AssertNotNil(err, t)
}

func TestRepository_UpdateHeadersAndCookiesNoCookiesTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsCookies(db)
	defer test_utils.DropTables(db)

	err := r.UpdateHeadersAndCookies(
		"bad-hash",
		[]models.UpdateModel{
			{
				Hash:      test_utils.HeaderHash1,
				FieldName: "key",
				NewValue:  "FOO",
			},
		},
		[]models.UpdateModel{
			{
				Hash:      test_utils.CookieHash1,
				FieldName: "value",
				NewValue:  "BAR",
			},
		},
	)

	test_utils.AssertNotNil(err, t)
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
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.DeleteHeader("bad-hash", test_utils.HeaderHash1)

	var headerExists bool
	_ = db.Get(
		&headerExists,
		`SELECT 1 FROM commands_headers WHERE hash = ?`,
		test_utils.HeaderHash1,
	)

	test_utils.AssertNotNil(err, t)
	test_utils.AssertTrue(headerExists, t)
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

func TestRepository_DeleteCookieBadAccountHash(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	err := r.DeleteCookie("bad-hash", test_utils.CookieHash1)

	var cookieExists bool
	_ = db.Get(
		&cookieExists,
		`SELECT 1 FROM commands_cookies WHERE hash = ?`,
		test_utils.CookieHash1,
	)

	test_utils.AssertNotNil(err, t)
	test_utils.AssertTrue(cookieExists, t)
}
