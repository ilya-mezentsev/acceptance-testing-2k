package session

import (
	"api_meta/mock/services"
	"api_meta/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"services/errors"
	"services/plugins/response_factory"
	"test_utils"
	"testing"
)

var (
	repository = services.SessionRepositoryMock{Accounts: map[string]bool{
		services.ExistsAccountHash: true,
	}}
	s                     = New(repository)
	expectedSuccessStatus = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus   = response_factory.ErrorResponse(nil).GetStatus()
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_CreateSessionSuccess(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request := test_utils.GetMockRequest(fmt.Sprintf(
		`{"login": "%s", "password": "%s"}`,
		services.ExistsLogin,
		services.ExistsPassword,
	))

	response := s.CreateSession(responseRecorder, request)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		services.ExistsAccountHash,
		response.GetData().(models.SessionResponse).AccountHash,
		t,
	)
	test_utils.AssertEqual(
		services.ExistsLogin,
		response.GetData().(models.SessionResponse).Login,
		t,
	)

	cookie := responseRecorder.Result().Cookies()[0]
	test_utils.AssertEqual(CookieName, cookie.Name, t)
	test_utils.AssertEqual(services.ExistsAccountHash, cookie.Value, t)
	test_utils.AssertEqual("/", cookie.Path, t)
	test_utils.AssertTrue(cookie.HttpOnly, t)
}

func TestService_CreateSessionDecodeBodyError(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request := test_utils.GetMockRequest(`1`)

	response := s.CreateSession(responseRecorder, request)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateSessionCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateSessionInvalidLogin(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request := test_utils.GetMockRequest(`{"login": "@#@$FF#", "password": "blah-blah"}`)

	response := s.CreateSession(responseRecorder, request)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateSessionCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateSessionBadAccountHash(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request := test_utils.GetMockRequest(fmt.Sprintf(
		`{"login": "%s", "password": "%s"}`,
		services.BadLogin,
		services.BadPassword,
	))

	response := s.CreateSession(responseRecorder, request)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateSessionCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateSessionAccountDoesNotExists(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request := test_utils.GetMockRequest(`{"login": "blah-blah", "password": "saf435"}`)

	response := s.CreateSession(responseRecorder, request)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateSessionCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		accountDoesNotExistsError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetSessionSuccess(t *testing.T) {
	request := test_utils.GetMockRequest(fmt.Sprintf(
		`{"login": "%s", "password": "%s"}`,
		services.ExistsLogin,
		services.BadPassword,
	))
	request.AddCookie(&http.Cookie{
		Name:  CookieName,
		Value: "some-hash",
	})

	response := s.GetSession(request)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		"some-hash",
		response.GetData().(models.SessionResponse).AccountHash,
		t,
	)
}

func TestService_GetSessionNoSessionCookieError(t *testing.T) {
	request := test_utils.GetMockRequest(fmt.Sprintf(
		`{"login": "%s", "password": "%s"}`,
		services.ExistsLogin,
		services.BadPassword,
	))

	response := s.GetSession(request)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToGetSessionCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		sessionCookieNotFoundError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_DeleteSession(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	http.SetCookie(responseRecorder, &http.Cookie{
		Name:     CookieName,
		Value:    "some-hash",
		Path:     "/",
		HttpOnly: true,
	})

	response := s.DeleteSession(responseRecorder)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)

	cookie := responseRecorder.Result().Cookies()[1]
	test_utils.AssertEqual(CookieName, cookie.Name, t)
	test_utils.AssertEqual("", cookie.Value, t)
	test_utils.AssertEqual("/", cookie.Path, t)
	test_utils.AssertTrue(cookie.HttpOnly, t)
}
