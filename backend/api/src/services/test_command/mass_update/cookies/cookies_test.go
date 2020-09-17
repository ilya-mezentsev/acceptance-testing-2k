package cookies

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
			"cookies": [{"key": "Some-cookie", "value": "some-value"}],
			"command_hashes": [{"hash": "%s"}, {"hash": "%s"}]
		}`,
			services.PredefinedCommandHash1,
			services.PredefinedCommandHash2),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	firstCreatedCookie := repository.Meta[services.PredefinedAccountHash][1].Cookies[0]
	secondCreatedCookie := repository.Meta[services.PredefinedAccountHash][1].Cookies[1]

	test_utils.AssertEqual("Some-cookie", firstCreatedCookie.Key, t)
	test_utils.AssertEqual("some-value", firstCreatedCookie.Value, t)
	test_utils.AssertEqual("Some-cookie", secondCreatedCookie.Key, t)
	test_utils.AssertEqual("some-value", secondCreatedCookie.Value, t)
}

func TestService_CreateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(``, test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToMassAddCookies,
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

	response := s.Create(services.PredefinedAccountHash, test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"cookies": [{"key": "++==", "value": "some-value"}],
			"command_hashes": [{"hash": "%s"}, {"hash": "%s"}]
		}`,
			services.PredefinedCommandHash1,
			services.PredefinedCommandHash2),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToMassAddCookies,
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
			"cookies": [{"key": "Some-cookie", "value": "some-value"}],
			"command_hashes": [{"hash": "%s"}, {"hash": "%s"}]
		}`,
			services.PredefinedCommandHash1,
			services.PredefinedCommandHash2),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToMassAddCookies,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
