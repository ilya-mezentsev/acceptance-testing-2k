package meta_creator

import (
	"api_meta/mock/services"
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
	repository            = services.TestCommandKeyValueCreatorRepositoryMock{}
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
		fmt.Sprintf(`{
			"account_hash": "%s",
			"command_hash": "%s",
			"command_meta": {
				"headers": [{"key": "X-Token", "value": "token"}],
				"cookies": [{"key": "CSRF-Token", "value": "csrf-token"}]
			}
		}`, services.PredefinedAccountHash, services.PredefinedTestCommand1.Hash),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	createdKeyValue := repository.KeyValues[services.PredefinedAccountHash][0]
	test_utils.AssertEqual("X-Token", createdKeyValue.Headers[0].Key, t)
	test_utils.AssertEqual("token", createdKeyValue.Headers[0].Value, t)
	test_utils.AssertEqual(
		services.PredefinedTestCommand1.Hash,
		createdKeyValue.Headers[0].CommandHash,
		t,
	)
	test_utils.AssertEqual("CSRF-Token", createdKeyValue.Cookies[0].Key, t)
	test_utils.AssertEqual("csrf-token", createdKeyValue.Cookies[0].Value, t)
	test_utils.AssertEqual(
		services.PredefinedTestCommand1.Hash,
		createdKeyValue.Cookies[0].CommandHash,
		t,
	)
}

func TestService_CreateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateCommandMeta,
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
		`{
			"account_hash": "bad-hash",
			"command_hash": "bad-hash",
			"command_meta": {
				"headers": [{"key": "!@#@$#@", "value": "token"}]
			}
		}`,
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateCommandMeta,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := New(
		services.TestCommandKeyValueCreatorErroredRepositoryMock{},
	).Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"account_hash": "%s",
			"command_hash": "%s",
			"command_meta": {
				"headers": [{"key": "X-Token", "value": "token"}],
				"cookies": [{"key": "CSRF-Token", "value": "csrf-token"}]
			}
		}`, services.PredefinedAccountHash, services.PredefinedTestCommand1.Hash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateCommandMeta,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
