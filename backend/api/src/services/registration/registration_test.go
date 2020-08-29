package registration

import (
	"api_meta/interfaces"
	"api_meta/mock/services"
	"env"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"services/errors"
	"services/plugins/account_credentials"
	"services/plugins/response_factory"
	"test_utils"
	"testing"
	"utils"
)

var (
	s                     interfaces.CreateService
	filesRootPath         string
	repository            = services.RegistrationRepositoryMock{}
	expectedSuccessStatus = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus   = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	filesRootPath = utils.MustGetEnv("REGISTRATION_ROOT_PATH")

	s = New(&repository, filesRootPath)
	repository.Reset()
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	_ = os.RemoveAll(filesRootPath)
	res := m.Run()
	_ = os.RemoveAll(filesRootPath)
	os.Exit(res)
}

func TestService_RegisterSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`{"login": "some-login", "password": "!@#@!%@#%"}`))

	expectedHash := account_credentials.GenerateAccountHash("some-login")
	_, hashCreated := repository.AccountHashes[expectedHash]
	_, credentialsCreated := repository.AccountCredentials[expectedHash]
	test_utils.AssertTrue(hashCreated, t)
	test_utils.AssertTrue(credentialsCreated, t)
	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
	test_utils.AssertTrue(
		test_utils.MustFileExists(path.Join(filesRootPath, expectedHash, env.DBFilename)),
		t,
	)
}

func TestService_RegisterDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`1`))
	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToRegisterAccountCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_RegisterLoginExists(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(fmt.Sprintf(
		`{"login": "%s", "password": "%s"}`,
		services.ExistsLogin, services.ExistsPassword,
	)))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(unableToRegisterAccountCode, response.GetData().(errors.ServiceError).Code, t)
	test_utils.AssertEqual(loginAlreadyExistsError, response.GetData().(errors.ServiceError).Description, t)
}

func TestService_RegisterInvalidLogin(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`{"login": "", "password": "!@#@!%@#%"}`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(unableToRegisterAccountCode, response.GetData().(errors.ServiceError).Code, t)
	test_utils.AssertEqual(errors.InvalidRequestError, response.GetData().(errors.ServiceError).Description, t)
}

func TestService_RegisterBadAccountHash(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(fmt.Sprintf(
		`{"login": "%s", "password": "%s"}`,
		services.BadLogin, services.BadPassword,
	)))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(unableToRegisterAccountCode, response.GetData().(errors.ServiceError).Code, t)
	test_utils.AssertEqual(errors.RepositoryError, response.GetData().(errors.ServiceError).Description, t)
}

func TestService_RegisterBadLogin(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(fmt.Sprintf(
		`{"login": "%s", "password": "some-password"}`, services.BadLogin,
	)))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(unableToRegisterAccountCode, response.GetData().(errors.ServiceError).Code, t)
	test_utils.AssertEqual(errors.RepositoryError, response.GetData().(errors.ServiceError).Description, t)
}

func TestService_RegisterCannotCreateDir(t *testing.T) {
	s := New(&repository, "/")

	response := s.Create(test_utils.GetReadCloser(`{"login": "some-login", "password": "!@#@!%@#%"}`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(unableToRegisterAccountCode, response.GetData().(errors.ServiceError).Code, t)
	test_utils.AssertEqual(errors.UnknownError, response.GetData().(errors.ServiceError).Description, t)
}

func TestService_RegisterCannotCreateDBFile(t *testing.T) {
	defer repository.Reset()

	expectedHash := account_credentials.GenerateAccountHash("some-login")
	_ = os.Chmod(path.Join(filesRootPath, expectedHash, env.DBFilename), 0100)

	response := s.Create(test_utils.GetReadCloser(`{"login": "some-login", "password": "!@#@!%@#%"}`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(unableToRegisterAccountCode, response.GetData().(errors.ServiceError).Code, t)
	test_utils.AssertEqual(errors.UnknownError, response.GetData().(errors.ServiceError).Description, t)
}
