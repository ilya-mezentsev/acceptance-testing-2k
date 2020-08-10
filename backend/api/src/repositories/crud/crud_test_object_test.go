package crud

import (
	"api_meta/interfaces"
	"api_meta/models"
	"db_connector"
	"github.com/jmoiron/sqlx"
	"path"
	"repositories/crud/query_providers/test_object"
	"test_utils"
	"testing"
	"utils"
)

var (
	testHash   string
	db         *sqlx.DB
	connector  db_connector.Connector
	repository interfaces.CRUDRepository
)

func init() {
	testHash = utils.MustGetEnv("TEST_ACCOUNT_HASH")
	connector = db_connector.New(path.Dir(utils.MustGetEnv("TEST_DB_FILE")))
	db, _ = connector.Connect(testHash)
	repository = New(connector, test_object.QueryProvider{})
}

func TestRepository_TestObject(t *testing.T) {
	for name, fn := range map[string]func(t *testing.T){
		"Create TestObject Success":        createTestObjectSuccess,
		"Create TestObject BadAccountHash": createTestObjectBadAccountHash,

		"GetAll TestObjects Success":        getAllTestObjectsSuccess,
		"GetAll TestObjects BadAccountHash": getAllTestObjectsBadAccountHash,

		"Get TestObject Success":        getTestObjectSuccess,
		"Get TestObject NotFound":       getTestObjectNotFound,
		"Get TestObject BadAccountHash": getTestObjectBadAccountHash,

		"Update TestObject Success":        updateTestObjectSuccess,
		"Update TestObject BadFieldName":   updateTestObjectBadFieldName,
		"Update TestObject BadAccountHash": updateTestObjectBadAccountHash,

		"Delete TestObject Success":        deleteTestObjectSuccess,
		"Delete TestObject BadAccountHash": deleteTestObjectBadAccountHash,
	} {
		test_utils.InitTables(db)
		t.Run(name, func(t *testing.T) {
			fn(t)
		})
	}
}

func createTestObjectSuccess(t *testing.T) {
	expectedObject := models.TestObject{
		Name: "SOME_OBJECT",
		Hash: "some-hash",
	}
	err := repository.Create(testHash, map[string]interface{}{
		"name": expectedObject.Name,
		"hash": expectedObject.Hash,
	})

	var createdObject models.TestObject
	_ = repository.Get(testHash, "some-hash", &createdObject)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedObject, createdObject, t)
}

func createTestObjectBadAccountHash(t *testing.T) {
	err := repository.Create(test_utils.NotExistsAccountHash, map[string]interface{}{
		"name": "SOME_NAME",
		"hash": "some-hash",
	})

	test_utils.AssertNotNil(err, t)
}

func getAllTestObjectsSuccess(t *testing.T) {
	var objects []models.TestObject
	err := repository.GetAll(testHash, &objects)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.ObjectName, objects[0].Name, t)
	test_utils.AssertEqual(test_utils.ObjectHash, objects[0].Hash, t)
}

func getAllTestObjectsBadAccountHash(t *testing.T) {
	var objects []models.TestObject
	err := repository.GetAll(test_utils.NotExistsAccountHash, &objects)

	test_utils.AssertNotNil(err, t)
}

func getTestObjectSuccess(t *testing.T) {
	var object models.TestObject
	err := repository.Get(testHash, test_utils.ObjectHash, &object)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.ObjectName, object.Name, t)
	test_utils.AssertEqual(test_utils.ObjectHash, object.Hash, t)
}

func getTestObjectNotFound(t *testing.T) {
	var object models.TestObject
	err := repository.Get(testHash, test_utils.NotExistsObjectHash, &object)

	test_utils.AssertNotNil(err, t)
}

func getTestObjectBadAccountHash(t *testing.T) {
	var object models.TestObject
	err := repository.Get(test_utils.NotExistsAccountHash, test_utils.ObjectHash, &object)

	test_utils.AssertNotNil(err, t)
}

func updateTestObjectSuccess(t *testing.T) {
	err := repository.Update(testHash, []models.UpdateModel{
		{
			Hash:      test_utils.ObjectHash,
			FieldName: "name",
			NewValue:  "FOO",
		},
	})

	var updatedObject models.TestObject
	_ = repository.Get(testHash, test_utils.ObjectHash, &updatedObject)
	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("FOO", updatedObject.Name, t)
}

func updateTestObjectBadFieldName(t *testing.T) {
	err := repository.Update(testHash, []models.UpdateModel{
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

	err := repository.Update(test_utils.NotExistsAccountHash, []models.UpdateModel{})

	test_utils.AssertNotNil(err, t)
}

func deleteTestObjectSuccess(t *testing.T) {
	err := repository.Delete(testHash, test_utils.ObjectHash)

	test_utils.AssertNil(err, t)
	var objects []models.TestObject
	err = repository.GetAll(testHash, &objects)
	test_utils.AssertEqual(0, len(objects), t)
}

func deleteTestObjectBadAccountHash(t *testing.T) {
	err := repository.Delete(test_utils.NotExistsAccountHash, test_utils.ObjectHash)

	test_utils.AssertNotNil(err, t)
}
