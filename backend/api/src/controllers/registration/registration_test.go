package registration

import (
	"api_meta/mock/services"
	"controllers/plugins/response_writer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"os"
	"services/errors"
	"services/plugins/response_factory"
	"services/registration"
	"test_utils"
	"testing"
	"utils"
)

var (
	r             *mux.Router
	filesRootPath string
	repository    = services.RegistrationRepositoryMock{}
)

func init() {
	filesRootPath = utils.MustGetEnv("REGISTRATION_ROOT_PATH")
	r = mux.NewRouter()
	s := registration.New(&repository, filesRootPath)

	repository.Reset()
	Init(r, s)
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	_ = os.RemoveAll(filesRootPath)
	res := m.Run()
	_ = os.RemoveAll(filesRootPath)
	os.Exit(res)
}

func Test_CreateAccountSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer repository.Reset()
	hashesCount := len(repository.AccountHashes)
	credentialsCount := len(repository.AccountCredentials)

	responseData := test_utils.RequestPost(fmt.Sprintf(
		"%s/registration/", server.URL,
	), `{"login": "some-login", "password": "!@#@!%@#%"}`)

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.DefaultResponse().GetStatus(), response.Status, t)
	test_utils.AssertEqual(hashesCount+1, len(repository.AccountHashes), t)
	test_utils.AssertEqual(credentialsCount+1, len(repository.AccountCredentials), t)
}

func Test_CreateAccountError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer repository.Reset()

	responseData := test_utils.RequestPost(fmt.Sprintf(
		"%s/registration/", server.URL,
	), `1`)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.ErrorResponse(nil).GetStatus(), response.Status, t)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.Data.(map[string]interface{})["description"],
		t,
	)
}
