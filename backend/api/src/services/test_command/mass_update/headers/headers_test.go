package headers

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

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"account_hash": "%s",
			"headers": [{"key": "Some-header", "value": "some-value"}],
			"command_hashes": [{"hash": "%s"}, {"hash": "%s"}]
		}`,
			services.PredefinedAccountHash,
			services.PredefinedCommandHash1,
			services.PredefinedCommandHash2),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	firstCreatedHeader := repository.Meta[services.PredefinedAccountHash][1].Headers[0]
	secondCreatedHeader := repository.Meta[services.PredefinedAccountHash][1].Headers[1]

	test_utils.AssertEqual("Some-header", firstCreatedHeader.Key, t)
	test_utils.AssertEqual("some-value", firstCreatedHeader.Value, t)
	test_utils.AssertEqual("Some-header", secondCreatedHeader.Key, t)
	test_utils.AssertEqual("some-value", secondCreatedHeader.Value, t)
}

func TestService_CreateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToMassAddHeaders,
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
		fmt.Sprintf(`{
			"account_hash": "%s",
			"headers": [{"key": "++==", "value": "some-value"}],
			"command_hashes": [{"hash": "%s"}, {"hash": "%s"}]
		}`,
			services.PredefinedAccountHash,
			services.PredefinedCommandHash1,
			services.PredefinedCommandHash2),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToMassAddHeaders,
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

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"account_hash": "%s",
			"headers": [{"key": "Some-cookie", "value": "some-value"}],
			"command_hashes": [{"hash": "%s"}, {"hash": "%s"}]
		}`,
			services.BadAccountHash,
			services.PredefinedCommandHash1,
			services.PredefinedCommandHash2),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToMassAddHeaders,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
