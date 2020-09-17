package meta

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
	repository            = services.TestCommandMetaRepositoryMock{}
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

	response := s.Create(services.PredefinedAccountHash, test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"command_hash": "%s",
			"command_meta": {
				"headers": [{"key": "X-Token", "value": "token"}],
				"cookies": [{"key": "CSRF-Token", "value": "csrf-token"}]
			}
		}`, services.PredefinedTestCommand1.Hash),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	allMeta := repository.Meta[services.PredefinedAccountHash]
	createdMeta := allMeta[len(allMeta)-1]
	test_utils.AssertEqual("X-Token", createdMeta.Headers[0].Key, t)
	test_utils.AssertEqual("token", createdMeta.Headers[0].Value, t)
	test_utils.AssertEqual(
		services.PredefinedTestCommand1.Hash,
		createdMeta.Headers[0].CommandHash,
		t,
	)
	test_utils.AssertEqual("CSRF-Token", createdMeta.Cookies[0].Key, t)
	test_utils.AssertEqual("csrf-token", createdMeta.Cookies[0].Value, t)
	test_utils.AssertEqual(
		services.PredefinedTestCommand1.Hash,
		createdMeta.Cookies[0].CommandHash,
		t,
	)
}

func TestService_CreateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(``, test_utils.GetReadCloser(`1`))

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

	response := s.Create(``, test_utils.GetReadCloser(
		`{
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

	response := s.Create(services.BadAccountHash, test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"command_hash": "%s",
			"command_meta": {
				"headers": [{"key": "X-Token", "value": "token"}],
				"cookies": [{"key": "CSRF-Token", "value": "csrf-token"}]
			}
		}`, services.PredefinedTestCommand1.Hash),
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

func TestService_UpdateSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Update(services.PredefinedAccountHash, test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"headers": [{"hash": "%s", "field_name": "key", "new_value": "FOO"}],
			"cookies": [{"hash": "%s", "field_name": "value", "new_value": "BAR"}]
		}`,
			services.PredefinedHeader1.Hash,
			services.PredefinedCookie1.Hash),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	var (
		cookieUpdated, headerUpdated bool
	)
	for _, meta := range repository.Meta[services.PredefinedAccountHash] {
		for _, header := range meta.Headers {
			if header.Hash == services.PredefinedHeader1.Hash {
				headerUpdated = header.Key == "FOO"
			}
		}

		for _, cookie := range meta.Cookies {
			if cookie.Hash == services.PredefinedCookie1.Hash {
				cookieUpdated = cookie.Value == "BAR"
			}
		}
	}

	test_utils.AssertTrue(cookieUpdated && headerUpdated, t)
}

func TestService_UpdateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(``, test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateCommandMeta,
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

	response := s.Update(services.PredefinedAccountHash, test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"headers": [{"hash": "%s", "field_name": "foo", "new_value": "FOO"}],
			"cookies": [{"hash": "%s", "field_name": "bar", "new_value": "BAR"}]
		}`,
			services.PredefinedHeader1.Hash,
			services.PredefinedCookie1.Hash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateCommandMeta,
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

	response := s.Update(services.BadAccountHash, test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"headers": [{"hash": "%s", "field_name": "key", "new_value": "FOO"}],
			"cookies": [{"hash": "%s", "field_name": "value", "new_value": "BAR"}]
		}`,
			services.PredefinedHeader1.Hash,
			services.PredefinedCookie1.Hash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateCommandMeta,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
