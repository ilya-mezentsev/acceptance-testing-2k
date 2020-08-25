package test_object

import (
	"api_meta/mock/services"
	"api_meta/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"services/errors"
	"services/plugins/response_factory"
	"test_utils"
	"testing"
)

var (
	repository            = services.TestObjectRepositoryMock{}
	s                     = New(&repository)
	expectedSuccessStatus = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus   = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	repository.Reset()
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_CreateSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{"account_hash": "%s", "test_object": {"name": "TEST"}}`, services.SomeHash),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
	test_utils.AssertNotNil(repository.Objects[services.SomeHash][0].Hash, t)
	test_utils.AssertEqual("TEST", repository.Objects[services.SomeHash][0].Name, t)
}

func TestService_CreateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		`{"account_hash": "some-hash", "test_object": {"name": "TEST"}}`,
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateObjectExistsError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(
			`{"account_hash": "%s", "test_object": {"name": "%s"}}`,
			services.PredefinedAccountHash, services.PredefinedTestObject1.Name,
		),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.UniqueEntityExistsError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(
			`{"account_hash": "%s", "test_object": {"name": "TEST"}}`,
			services.BadAccountHash,
		),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetAllSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.GetAll(services.PredefinedAccountHash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	for expectedObjectIndex, expectedObject := range repository.Objects[services.PredefinedAccountHash] {
		test_utils.AssertEqual(
			expectedObject,
			response.GetData().([]models.TestObject)[expectedObjectIndex],
			t,
		)
	}
}

func TestService_GetAllInvalidHashError(t *testing.T) {
	defer repository.Reset()

	response := s.GetAll("some-hash")
	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToFetchTestObjectsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetAllRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.GetAll(services.BadAccountHash)
	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToFetchTestObjectsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Get(services.PredefinedAccountHash, services.PredefinedTestObject1.Hash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		services.PredefinedTestObject1.Name,
		response.GetData().(models.TestObject).Name,
		t,
	)
	test_utils.AssertEqual(
		services.PredefinedTestObject1.Hash,
		response.GetData().(models.TestObject).Hash,
		t,
	)
}

func TestService_GetInvalidHashError(t *testing.T) {
	defer repository.Reset()

	response := s.Get("hash-1", "hash-2")

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToFetchTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Get(services.BadAccountHash, services.PredefinedTestObject1.Hash)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToFetchTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_UpdateSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(
		fmt.Sprintf(
			`
			{"account_hash": "%s",
			"update_payload": [{"hash": "%s", "field_name": "name", "new_value": "FOO"}]
			}`,
			services.PredefinedAccountHash,
			services.PredefinedTestObject1.Hash,
		),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
	test_utils.AssertEqual("FOO", repository.Objects[services.PredefinedAccountHash][0].Name, t)
}

func TestService_UpdateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_UpdateInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(
		`{
			"account_hash": "hash-1",
			"update_payload": [{"hash": "hash-2", "field_name": "bad-name", "new_value": "FOO"}]
		}`,
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_UpdateRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(
		fmt.Sprintf(
			`
			{"account_hash": "%s",
			"update_payload": [{"hash": "%s", "field_name": "name", "new_value": "FOO"}]
			}`,
			services.BadAccountHash,
			services.PredefinedTestObject1.Hash,
		),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_DeleteSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Delete(services.PredefinedAccountHash, services.PredefinedTestObject1.Hash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
	for _, object := range repository.Objects[services.PredefinedAccountHash] {
		test_utils.AssertNotEqual(
			services.PredefinedTestObject1.Hash,
			object.Hash,
			t,
		)
	}
}

func TestService_DeleteInvalidHashError(t *testing.T) {
	defer repository.Reset()

	response := s.Delete("hash-1", "hash-2")

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToDeleteTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_DeleteRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Delete(services.BadAccountHash, services.PredefinedTestObject1.Hash)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToDeleteTestObjectCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
