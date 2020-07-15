package response_parser

import (
	"command/http/errors"
	"io/ioutil"
	"log"
	mockCommand "mock/command"
	"net/http"
	"os"
	"strings"
	"testing"
	"utils"
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

	utils.AssertNil(err, t)
	utils.AssertEqual("1", data["x"], t)
	utils.AssertEqual("true", data["y"], t)
	utils.AssertEqual("hey", data["z"], t)
}

func TestParseDataWithArray(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": [0, 1, 2, 3, 4.4, 5.05]}`)),
	})
	expectedSlice := []string{"0", "1", "2", "3", "4.4", "5.05"}
	currentSlice, ok := data["x"].([]interface{})

	utils.AssertNil(err, t)
	utils.AssertTrue(ok, t)
	for i, expectedValue := range expectedSlice {
		utils.AssertEqual(expectedValue, currentSlice[i], t)
	}
}

func TestParseDataWithMap(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": {"y": 1}}`)),
	})
	currentMap, ok := data["x"].(map[string]interface{})

	utils.AssertNil(err, t)
	utils.AssertTrue(ok, t)
	utils.AssertEqual("1", currentMap["y"], t)
}

func TestParseInvalidJSON(t *testing.T) {
	_, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(``)),
	})

	utils.AssertErrorsEqual(errors.NoJSONInResponse, err, t)
}

func TestParseBrokenResponseBody(t *testing.T) {
	_, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       mockCommand.BadReadCloser{},
	})

	utils.AssertErrorsEqual(errors.NoJSONInResponse, err, t)
}

func TestProcessValueUnknownType(t *testing.T) {
	value := processValue(1)

	utils.AssertNil(value, t)
}
