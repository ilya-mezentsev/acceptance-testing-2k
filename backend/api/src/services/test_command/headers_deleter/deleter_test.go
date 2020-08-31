package headers_deleter

import (
	"api_meta/mock/services"
	"io/ioutil"
	"log"
	"os"
	"services/errors"
	"services/plugins/response_factory"
	"test_utils"
	"testing"
)

var (
	repository            = services.TestCommandHeadersDeleterRepositoryMock{}
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

func TestService_DeleteSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Delete(services.PredefinedAccountHash, services.PredefinedHeader1.Hash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	for _, header := range repository.Headers[services.PredefinedAccountHash] {
		test_utils.AssertNotEqual(services.PredefinedHeader1.Hash, header.Hash, t)
	}
}

func TestService_DeleteInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Delete("bad-hash", "bad-hash")

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToDeleteCommandHeader,
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

	response := s.Delete(services.BadAccountHash, services.PredefinedHeader1.Hash)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToDeleteCommandHeader,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
