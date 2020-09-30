package test_utils

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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

func RequestGet(url, accountHash string) io.ReadCloser {
	return makeRequest(http.MethodGet, url, "", accountHash)
}

func RequestPost(url, data, accountHash string) io.ReadCloser {
	return makeRequest(http.MethodPost, url, data, accountHash)
}

func RequestPatch(url, data, accountHash string) io.ReadCloser {
	return makeRequest(http.MethodPatch, url, data, accountHash)
}

func RequestDelete(url, accountHash string) io.ReadCloser {
	return makeRequest(http.MethodDelete, url, "", accountHash)
}

func makeRequest(method, url, data, accountHash string) io.ReadCloser {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "AT2K-Session",
		Value: accountHash,
	})
	return MustDoRequest(req)
}

func MustDoRequest(req *http.Request) io.ReadCloser {
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}

	return res.Body
}

func GetReadCloser(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}

func MustFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func GetMockRequest(data string) *http.Request {
	return httptest.NewRequest(
		http.MethodPost,
		"https://link.com",
		bytes.NewBuffer([]byte(data)),
	)
}

func MustGetFileUploadRequest(url, paramName, filePath, accountHash string) *http.Request {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fileInfo.Name())
	if err != nil {
		panic(err)
	}

	_, err = part.Write(fileContents)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req
}

func MustGetFileUploadMockRequest(paramName, filePath, accountHash string) *http.Request {
	return MustGetFileUploadRequest("https://link.com", paramName, filePath, accountHash)
}
