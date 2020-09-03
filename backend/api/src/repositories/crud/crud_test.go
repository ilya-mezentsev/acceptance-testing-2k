package crud

import (
	"api_meta/interfaces"
	"api_meta/models"
	"api_meta/types"
	"db_connector"
	"github.com/jmoiron/sqlx"
	"path"
	"repositories/crud/query_providers"
	"services/plugins/hash"
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
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)

	testObjectRepository = New(connector, query_providers.TestObjectQueryProvider{})
	testCommandRepository = New(connector, query_providers.TestCommandQueryProvider{})
}

func TestRepository(t *testing.T) {
	for name, fn := range map[string]func(t *testing.T){
		"Create TestObject Success":           createTestObjectSuccess,
		"Create TestObject BadAccountHash":    createTestObjectBadAccountHash,
		"Create TestObject NameAlreadyExists": createTestObjectWithExistsName,

		"Create CommandSettings Success": createTestCommandSuccess,

		"GetAll TestObjects Success":        getAllTestObjectsSuccess,
		"GetAll TestObjects BadAccountHash": getAllTestObjectsBadAccountHash,

		"GetAll TestCommands Success": getAllTestCommandsSuccess,

		"Get TestObject Success":        getTestObjectSuccess,
		"Get TestObject NotFound":       getTestObjectNotFound,
		"Get TestObject BadAccountHash": getTestObjectBadAccountHash,

		"Get CommandSettings Success": getTestCommandSuccess,

		"Update TestObject Success":        updateTestObjectSuccess,
		"Update TestObject BadFieldName":   updateTestObjectBadFieldName,
		"Update TestObject BadAccountHash": updateTestObjectBadAccountHash,

		"Update CommandSettings Success":         updateTestCommandSuccess,
		"Update CommandSettings BadUpdateTarget": updateTestObjectBadUpdateTarget,

		"Delete TestObject Success":        deleteTestObjectSuccess,
		"Delete TestObject BadAccountHash": deleteTestObjectBadAccountHash,

		"Delete CommandSettings Success": deleteTestCommandSuccess,
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
		Hash: hash.Md5WithTimeAsKey("some-hash"),
	}
	err := testObjectRepository.Create(testHash, map[string]interface{}{
		"name": expectedObject.Name,
		"hash": expectedObject.Hash,
	})

	var createdObject models.TestObject
	_ = testObjectRepository.Get(testHash, expectedObject.Hash, &createdObject)
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

func createTestObjectWithExistsName(t *testing.T) {
	err := testObjectRepository.Create(testHash, map[string]interface{}{
		"name": test_utils.ObjectName,
		"hash": "some-hash",
	})

	test_utils.AssertErrorsEqual(types.UniqueEntityAlreadyExists{}, err, t)
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
	expectedCommand := models.CommandSettings{
		Name:       "FOO",
		Hash:       hash.Md5WithTimeAsKey("some-hash"),
		ObjectHash: test_utils.ObjectHash,
		Method:     "GET",
		BaseURL:    "https://link.com",
		Endpoint:   "api/v2/user",
	}
	err := testCommandRepository.Create(testHash, map[string]interface{}{
		"name":                  expectedCommand.Name,
		"hash":                  expectedCommand.Hash,
		"object_hash":           expectedCommand.ObjectHash,
		"method":                expectedCommand.Method,
		"base_url":              expectedCommand.BaseURL,
		"endpoint":              expectedCommand.Endpoint,
		"pass_arguments_in_url": expectedCommand.PassArgumentsInURL,
	})

	var createdCommand models.CommandSettings
	_ = testCommandRepository.Get(testHash, expectedCommand.Hash, &createdCommand)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedCommand.Name, createdCommand.Name, t)
	test_utils.AssertEqual(expectedCommand.Hash, createdCommand.Hash, t)
	test_utils.AssertEqual(expectedCommand.ObjectHash, createdCommand.ObjectHash, t)
	test_utils.AssertEqual(expectedCommand.Method, createdCommand.Method, t)
	test_utils.AssertEqual(expectedCommand.BaseURL, createdCommand.BaseURL, t)
	test_utils.AssertEqual(expectedCommand.Endpoint, createdCommand.Endpoint, t)
	test_utils.AssertEqual(expectedCommand.PassArgumentsInURL, createdCommand.PassArgumentsInURL, t)
}

func getAllTestCommandsSuccess(t *testing.T) {
	var commands []models.CommandSettings
	err := testCommandRepository.GetAll(testHash, &commands)

	var expectedCommand models.CommandSettings
	_ = testCommandRepository.Get(testHash, test_utils.CreateCommandHash, &expectedCommand)
	test_utils.AssertNil(err, t)
	for _, command := range commands {
		if command.Hash == expectedCommand.Hash {
			test_utils.AssertEqual(expectedCommand, command, t)
		}
	}
}

func getTestCommandSuccess(t *testing.T) {
	var command models.CommandSettings
	err := testCommandRepository.Get(testHash, test_utils.CreateCommandHash, &command)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.CreateCommandHash, command.Hash, t)
	test_utils.AssertEqual(test_utils.CreateCommandName, command.Name, t)
	test_utils.AssertEqual(test_utils.ObjectHash, command.ObjectHash, t)
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
	})

	var updatedCommand models.CommandSettings
	_ = testCommandRepository.Get(testHash, test_utils.CreateCommandHash, &updatedCommand)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("FOO", updatedCommand.Name, t)
	test_utils.AssertEqual("api/v2/foo", updatedCommand.Endpoint, t)
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

	var headersCount int
	_ = db.Get(
		&headersCount,
		`SELECT count(*) FROM commands_headers WHERE command_hash = ?`,
		test_utils.CreateCommandHash,
	)
	test_utils.AssertEqual(0, headersCount, t)

	var commands []models.CommandSettings
	_ = testCommandRepository.GetAll(testHash, &commands)
	test_utils.AssertNil(err, t)
	for _, command := range commands {
		test_utils.AssertNotEqual(test_utils.CreateCommandHash, command.Hash, t)
	}
}
