package response_parser

import (
	"command/http/errors"
	"io/ioutil"
	"log"
	mockCommand "mock/command"
	"net/http"
	"os"
	"strings"
	"test_utils"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestParseSimpleFlatData(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": 1, "y": true, "z": "hey"}`)),
	})

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual("1", data["x"], t)
	test_utils.AssertEqual("true", data["y"], t)
	test_utils.AssertEqual("hey", data["z"], t)
}

func TestParseDataWithArray(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": [0, 1, 2, 3, 4.4, 5.05]}`)),
	})
	expectedSlice := []string{"0", "1", "2", "3", "4.4", "5.05"}
	currentSlice, ok := data["x"].([]interface{})

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(ok, t)
	for i, expectedValue := range expectedSlice {
		test_utils.AssertEqual(expectedValue, currentSlice[i], t)
	}
}

func TestParseDataWithMap(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": {"y": 1}}`)),
	})
	currentMap, ok := data["x"].(map[string]interface{})

	test_utils.AssertNil(err, t)
	test_utils.AssertTrue(ok, t)
	test_utils.AssertEqual("1", currentMap["y"], t)
}

func TestParseInvalidJSON(t *testing.T) {
	_, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(``)),
	})

	test_utils.AssertErrorsEqual(errors.NoJSONInResponse, err, t)
}

func TestParseBrokenResponseBody(t *testing.T) {
	_, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       mockCommand.BadReadCloser{},
	})

	test_utils.AssertErrorsEqual(errors.NoJSONInResponse, err, t)
}

func TestProcessValueUnknownType(t *testing.T) {
	value := processValue(1)

	test_utils.AssertNil(value, t)
}
