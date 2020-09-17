package session

import (
	"api_meta/mock/services"
	"controllers/plugins/response_writer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"services/plugins/response_factory"
	sessionService "services/session"
	"test_utils"
	"testing"
)

var (
	r          *mux.Router
	repository = services.SessionRepositoryMock{Accounts: map[string]bool{
		services.ExistsAccountHash: true,
	}}
	s                     = sessionService.New(repository)
	expectedSuccessStatus = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus   = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	r = mux.NewRouter()

	Init(r, s)
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestController_GetSessionSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()

	request, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/session/", server.URL),
		nil,
	)
	request.AddCookie(&http.Cookie{
		Name:  sessionService.CookieName,
		Value: "some-hash",
	})

	responseData := test_utils.MustDoRequest(request)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedSuccessStatus, response.Status, t)
	test_utils.AssertEqual(
		"some-hash",
		response.Data.(map[string]interface{})["account_hash"],
		t,
	)
}

func TestController_GetSessionError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()

	request, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/session/", server.URL),
		nil,
	)

	responseData := test_utils.MustDoRequest(request)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedErrorStatus, response.Status, t)
}

func TestController_CreateSessionSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()

	responseData := test_utils.RequestPost(
		fmt.Sprintf("%s/session/", server.URL),
		fmt.Sprintf(
			`{"login": "%s", "password": "%s"}`,
			services.ExistsLogin,
			services.ExistsPassword,
		),
		``,
	)

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedSuccessStatus, response.Status, t)
}

func TestController_CreateSessionError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()

	responseData := test_utils.RequestPost(
		fmt.Sprintf("%s/session/", server.URL),
		fmt.Sprintf(
			`{"login": "%s", "password": "%s"}`,
			services.BadLogin,
			services.BadPassword,
		),
		``,
	)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedErrorStatus, response.Status, t)
}

func TestController_DeleteSession(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()

	responseData := test_utils.RequestDelete(
		fmt.Sprintf("%s/session/", server.URL),
		``,
	)

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(expectedSuccessStatus, response.Status, t)
}
