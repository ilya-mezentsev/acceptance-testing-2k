package test_utils

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Expectation struct {
	Expected, Actual interface{}
}

func AssertTrue(actual bool, t *testing.T) {
	if !actual {
		logExpectationAndFail(true, actual, t)
	}
}

func AssertFalse(actual bool, t *testing.T) {
	if actual {
		logExpectationAndFail(false, actual, t)
	}
}

func AssertEqual(expected, actual interface{}, t *testing.T) {
	if expected != actual {
		logExpectationAndFail(expected, actual, t)
	}
}

func AssertNotEqual(notExpected, actual interface{}, t *testing.T) {
	if notExpected == actual {
		logExpectationAndFail(fmt.Sprintf("not equal to %v", notExpected), actual, t)
	}
}

func AssertNil(v interface{}, t *testing.T) {
	if v != nil {
		logExpectationAndFail(nil, v, t)
	}
}

func AssertNotNil(v interface{}, t *testing.T) {
	if v == nil {
		logExpectationAndFail("not nil", v, t)
	}
}

func AssertErrorsEqual(expectedErr, actualErr error, t *testing.T) {
	if expectedErr != actualErr {
		logExpectationAndFail(expectedErr, actualErr, t)
	}
}

func logExpectationAndFail(expected, actual interface{}, t *testing.T) {
	t.Log(
		GetExpectationString(
			Expectation{Expected: expected, Actual: actual}))
	t.Fail()
}

func GetExpectationString(e Expectation) string {
	return fmt.Sprintf("expected: %v, got: %v\n", e.Expected, e.Actual)
}

func GetTestServer(r *mux.Router) *httptest.Server {
	return httptest.NewServer(r)
}

func RequestGet(url string) io.ReadCloser {
	return request(http.MethodGet, url, "")
}

func RequestPost(url, data string) io.ReadCloser {
	return request(http.MethodPost, url, data)
}

func RequestPatch(url, data string) io.ReadCloser {
	return request(http.MethodPatch, url, data)
}

func RequestDelete(url string) io.ReadCloser {
	return request(http.MethodDelete, url, "")
}

func request(method, url, data string) io.ReadCloser {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}

	return res.Body
}

func GetReadCloser(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}
