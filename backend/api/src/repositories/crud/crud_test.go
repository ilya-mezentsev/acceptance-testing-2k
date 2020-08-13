package crud

import (
	"api_meta/interfaces"
	"api_meta/models"
	"database/sql"
	"db_connector"
	"fmt"
	"github.com/jmoiron/sqlx"
	"path"
	"repositories/crud/query_providers"
	"strings"
	"test_utils"
	"testing"
	"utils"
)

var (
	testHash              string
	db                    *sqlx.DB
	connector             db_connector.Connector
	testObjectRepository  interfaces.CRUDRepository
	testCommandRepository interfaces.CRUDRepository
	credentialsRepository interfaces.CRUDRepository
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)

	testObjectRepository = New(connector, query_providers.TestObjectQueryProvider{})
	testCommandRepository = New(connector, query_providers.TestCommandQueryProvider{})
	credentialsRepository = New(connector, query_providers.AccountCredentialsQueryProvider{})
}

func TestRepository(t *testing.T) {
	for name, fn := range map[string]func(t *testing.T){
		"Create TestObject Success":        createTestObjectSuccess,
		"Create TestObject BadAccountHash": createTestObjectBadAccountHash,

		"Create TestCommand Success": createTestCommandSuccess,

		"Create AccountCredentials Success": createAccountCredentialsSuccess,

		"GetAll TestObjects Success":        getAllTestObjectsSuccess,
		"GetAll TestObjects BadAccountHash": getAllTestObjectsBadAccountHash,

		"GetAll TestCommands Success": getAllTestCommandsSuccess,

		"GetAll AccountCredentials Error": getAllAccountCredentialsError,

		"Get TestObject Success":        getTestObjectSuccess,
		"Get TestObject NotFound":       getTestObjectNotFound,
		"Get TestObject BadAccountHash": getTestObjectBadAccountHash,

		"Get TestCommand Success": getTestCommandSuccess,

		"Get AccountCredentials Success": getAccountCredentialsSuccess,

		"Update TestObject Success":        updateTestObjectSuccess,
		"Update TestObject BadFieldName":   updateTestObjectBadFieldName,
		"Update TestObject BadAccountHash": updateTestObjectBadAccountHash,

		"Update TestCommand Success":         updateTestCommandSuccess,
		"Update TestCommand BadUpdateTarget": updateTestObjectBadUpdateTarget,

		"Update AccountCredentials Success": updateAccountCredentialsSuccess,

		"Delete TestObject Success":        deleteTestObjectSuccess,
		"Delete TestObject BadAccountHash": deleteTestObjectBadAccountHash,

		"Delete TestCommand Success": deleteTestCommandSuccess,

		"Delete AccountCredentials Success": deleteAccountCredentialsSuccess,
	} {
		t.Run(name, func(t *testing.T) {
			test_utils.InitTables(db)
			fn(t)
		})
	}
}

func createTestObjectSuccess(t *testing.T) {
	expectedObject := models.TestObject{
		Name: "SOME_OBJECT",
		Hash: "some-hash",
	}
	err := testObjectRepository.Create(testHash, map[string]interface{}{
		"name": expectedObject.Name,
		"hash": expectedObject.Hash,
	})

	var createdObject models.TestObject
	_ = testObjectRepository.Get(testHash, "some-hash", &createdObject)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedObject, createdObject, t)
}

func createTestObjectBadAccountHash(t *testing.T) {
	err := testObjectRepository.Create(test_utils.NotExistsAccountHash, map[string]interface{}{
		"name": "SOME_NAME",
		"hash": "some-hash",
	})

	test_utils.AssertNotNil(err, t)
}

func getAllTestObjectsSuccess(t *testing.T) {
	var objects []models.TestObject
	err := testObjectRepository.GetAll(testHash, &objects)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.ObjectName, objects[0].Name, t)
	test_utils.AssertEqual(test_utils.ObjectHash, objects[0].Hash, t)
}

func getAllTestObjectsBadAccountHash(t *testing.T) {
	var objects []models.TestObject
	err := testObjectRepository.GetAll(test_utils.NotExistsAccountHash, &objects)

	test_utils.AssertNotNil(err, t)
}

func getTestObjectSuccess(t *testing.T) {
	var object models.TestObject
	err := testObjectRepository.Get(testHash, test_utils.ObjectHash, &object)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.ObjectName, object.Name, t)
	test_utils.AssertEqual(test_utils.ObjectHash, object.Hash, t)
}

func getTestObjectNotFound(t *testing.T) {
	var object models.TestObject
	err := testObjectRepository.Get(testHash, test_utils.NotExistsObjectHash, &object)

	test_utils.AssertNotNil(err, t)
}

func getTestObjectBadAccountHash(t *testing.T) {
	var object models.TestObject
	err := testObjectRepository.Get(test_utils.NotExistsAccountHash, test_utils.ObjectHash, &object)

	test_utils.AssertNotNil(err, t)
}

func updateTestObjectSuccess(t *testing.T) {
	err := testObjectRepository.Update(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.ObjectHash,
			FieldName: "name",
			NewValue:  "FOO",
		},
	})

	var updatedObject models.TestObject
	_ = testObjectRepository.Get(testHash, test_utils.ObjectHash, &updatedObject)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("FOO", updatedObject.Name, t)
}

func updateTestObjectBadFieldName(t *testing.T) {
	err := testObjectRepository.Update(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.ObjectHash,
			FieldName: "blah-blah",
			NewValue:  "FOO",
		},
	})

	test_utils.AssertNotNil(err, t)
}

func updateTestObjectBadAccountHash(t *testing.T) {
	test_utils.DropTables(db)

	err := testObjectRepository.Update(test_utils.NotExistsAccountHash, []models.UpdateModel{})

	test_utils.AssertNotNil(err, t)
}

func deleteTestObjectSuccess(t *testing.T) {
	err := testObjectRepository.Delete(testHash, test_utils.ObjectHash)

	test_utils.AssertNil(err, t)
	var objects []models.TestObject
	err = testObjectRepository.GetAll(testHash, &objects)
	test_utils.AssertEqual(0, len(objects), t)
}

func deleteTestObjectBadAccountHash(t *testing.T) {
	err := testObjectRepository.Delete(test_utils.NotExistsAccountHash, test_utils.ObjectHash)

	test_utils.AssertNotNil(err, t)
}

func createTestCommandSuccess(t *testing.T) {
	expectedCommand := models.TestCommandRequest{
		CommandSettings: models.CommandSettings{
			Name:       "FOO",
			Hash:       "some-hash",
			ObjectName: test_utils.ObjectName,
			Method:     "GET",
			BaseURL:    "https://link.com",
			Endpoint:   "api/v2/user",
		},
		Headers: map[string]string{
			"X-Test": "header-value",
		},
		Cookies: map[string]string{
			"Test": "cookie-value",
		},
	}
	err := testCommandRepository.Create(testHash, map[string]interface{}{
		"name":                  expectedCommand.Name,
		"hash":                  expectedCommand.Hash,
		"object_name":           expectedCommand.ObjectName,
		"method":                expectedCommand.Method,
		"base_url":              expectedCommand.BaseURL,
		"endpoint":              expectedCommand.Endpoint,
		"pass_arguments_in_url": expectedCommand.PassArgumentsInURL,
		"command_headers":       expectedCommand.Headers.ReduceToRecordable(),
		"command_cookies":       expectedCommand.Cookies.ReduceToRecordable(),
	})

	var createdCommand models.TestCommandRecord
	_ = testCommandRepository.Get(testHash, expectedCommand.Hash, &createdCommand)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedCommand.Name, createdCommand.Name, t)
	test_utils.AssertEqual(expectedCommand.Hash, createdCommand.Hash, t)
	test_utils.AssertEqual(expectedCommand.ObjectName, createdCommand.ObjectName, t)
	test_utils.AssertEqual(expectedCommand.Method, createdCommand.Method, t)
	test_utils.AssertEqual(expectedCommand.BaseURL, createdCommand.BaseURL, t)
	test_utils.AssertEqual(expectedCommand.Endpoint, createdCommand.Endpoint, t)
	test_utils.AssertEqual(expectedCommand.PassArgumentsInURL, createdCommand.PassArgumentsInURL, t)
	test_utils.AssertEqual(expectedCommand.Headers.ReduceToRecordable(), createdCommand.Headers, t)
	test_utils.AssertEqual(expectedCommand.Cookies.ReduceToRecordable(), createdCommand.Cookies, t)
}

func getAllTestCommandsSuccess(t *testing.T) {
	var commands []models.TestCommandRecord
	err := testCommandRepository.GetAll(testHash, &commands)

	var expectedCommand models.TestCommandRecord
	_ = testCommandRepository.Get(testHash, test_utils.CreateCommandHash, &expectedCommand)
	test_utils.AssertNil(err, t)
	for _, command := range commands {
		if command.Hash == expectedCommand.Hash {
			test_utils.AssertEqual(expectedCommand, command, t)
		}
	}
}

func getTestCommandSuccess(t *testing.T) {
	var command models.TestCommandRecord
	err := testCommandRepository.Get(testHash, test_utils.CreateCommandHash, &command)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.CreateCommandHash, command.Hash, t)
	test_utils.AssertEqual(test_utils.CreateCommandName, command.Name, t)
	test_utils.AssertEqual(test_utils.ObjectName, command.ObjectName, t)
	var foundSettings bool
	for _, settings := range test_utils.Settings {
		if settings["command_hash"].(string) == command.Hash {
			foundSettings = true
			test_utils.AssertEqual(settings["method"].(string), command.Method, t)
			test_utils.AssertEqual(settings["base_url"].(string), command.BaseURL, t)
			test_utils.AssertEqual(settings["endpoint"].(string), command.Endpoint, t)
			test_utils.AssertEqual(settings["pass_arguments_in_url"], command.PassArgumentsInURL, t)
		}
	}
	test_utils.AssertTrue(foundSettings, t)

	var headersFound bool
	for _, headers := range test_utils.Headers {
		if headers["command_hash"].(string) == command.Hash {
			headersFound = true
			test_utils.AssertTrue(
				strings.Contains(
					command.Headers,
					fmt.Sprintf("%s=%s", headers["key"].(string), headers["value"].(string)),
				),
				t,
			)
		}
	}
	test_utils.AssertTrue(headersFound, t)

	var cookiesFound bool
	for _, cookies := range test_utils.Cookies {
		if cookies["command_hash"].(string) == command.Hash {
			cookiesFound = true
			test_utils.AssertTrue(
				strings.Contains(
					command.Cookies,
					fmt.Sprintf("%s=%s", cookies["key"].(string), cookies["value"].(string)),
				),
				t,
			)
		}
	}
	test_utils.AssertTrue(cookiesFound, t)
}

func updateTestCommandSuccess(t *testing.T) {
	err := testCommandRepository.Update(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.CreateCommandHash,
			FieldName: "command:name",
			NewValue:  "FOO",
		},
		{
			Hash:      test_utils.CreateCommandHash,
			FieldName: "command_setting:endpoint",
			NewValue:  "api/v2/foo",
		},
		{
			Hash:      test_utils.CreateCommandHash,
			FieldName: "command_header:key",
			NewValue:  "X-Foo",
		},
		{
			Hash:      test_utils.CreateCommandHash,
			FieldName: "command_cookie:value",
			NewValue:  "test-foo",
		},
	})

	var updatedCommand models.TestCommandRecord
	_ = testCommandRepository.Get(testHash, test_utils.CreateCommandHash, &updatedCommand)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("FOO", updatedCommand.Name, t)
	test_utils.AssertEqual("api/v2/foo", updatedCommand.Endpoint, t)
	test_utils.AssertTrue(strings.Contains(updatedCommand.Headers, "X-Foo="), t)
	test_utils.AssertTrue(strings.Contains(updatedCommand.Cookies, "=test-foo"), t)
}

func updateTestObjectBadUpdateTarget(t *testing.T) {
	defer func() {
		test_utils.AssertEqual("invalid update field name format", recover(), t)
	}()

	_ = testCommandRepository.Update(testHash, []models.UpdateModel{
		{
			Hash:      testHash,
			FieldName: "foo",
			NewValue:  nil,
		},
	})
	test_utils.AssertTrue(false, t)
}

func deleteTestCommandSuccess(t *testing.T) {
	err := testCommandRepository.Delete(testHash, test_utils.CreateCommandHash)

	var commands []models.TestCommandRecord
	_ = testCommandRepository.GetAll(testHash, &commands)
	test_utils.AssertNil(err, t)
	for _, command := range commands {
		test_utils.AssertNotEqual(test_utils.CreateCommandHash, command.Hash, t)
	}
}

func createAccountCredentialsSuccess(t *testing.T) {
	err := credentialsRepository.Create(testHash, map[string]interface{}{
		"login":    "login",
		"password": "some_password",
		"hash":     "hash-1",
	})

	var createdAccountCredentials models.AccountCredentials
	_ = credentialsRepository.Get(testHash, "hash-1", &createdAccountCredentials)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("login", createdAccountCredentials.Login, t)
	test_utils.AssertEqual("hash-1", createdAccountCredentials.Hash, t)
	test_utils.AssertFalse(createdAccountCredentials.Verified, t)
}

func getAllAccountCredentialsError(t *testing.T) {
	defer func() {
		test_utils.AssertEqual("should not used here", recover(), t)
	}()

	var i interface{}
	_ = credentialsRepository.GetAll(testHash, &i)
	test_utils.AssertTrue(false, t)
}

func getAccountCredentialsSuccess(t *testing.T) {
	var accountCredentials models.AccountCredentials
	err := credentialsRepository.Get(testHash, test_utils.CredentialsHash, &accountCredentials)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.CredentialsHash, accountCredentials.Hash, t)
	test_utils.AssertEqual(test_utils.CredentialsLogin, accountCredentials.Login, t)
	test_utils.AssertFalse(accountCredentials.Verified, t)
}

func updateAccountCredentialsSuccess(t *testing.T) {
	// can update only verified credentials
	_, _ = db.Exec(
		`UPDATE account_credentials SET verified = 1 WHERE hash = ?`,
		test_utils.CredentialsHash,
	)

	err := credentialsRepository.Update(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.CredentialsHash,
			FieldName: "login",
			NewValue:  "foo",
		},
		{
			Hash:      test_utils.CredentialsHash,
			FieldName: "verified",
			NewValue:  true,
		},
	})

	var updatedAccountCredentials models.AccountCredentials
	_ = credentialsRepository.Get(testHash, test_utils.CredentialsHash, &updatedAccountCredentials)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("foo", updatedAccountCredentials.Login, t)
	test_utils.AssertTrue(updatedAccountCredentials.Verified, t)
}

func deleteAccountCredentialsSuccess(t *testing.T) {
	_ = credentialsRepository.Create(testHash, map[string]interface{}{
		"login":    "login",
		"password": "some_password",
		"hash":     "hash-1",
	})

	err := credentialsRepository.Delete(testHash, "hash-1")

	test_utils.AssertNil(err, t)
	var c models.AccountCredentials
	err = db.Get(
		&c,
		`SELECT hash, login, verified FROM account_credentials WHERE hash = ?`,
		"hash-1",
	)
	test_utils.AssertErrorsEqual(sql.ErrNoRows, err, t)
}
