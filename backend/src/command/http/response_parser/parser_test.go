package response_parser

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"utils"
)

func TestParseSimpleFlatData(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": 1, "y": true}`)),
	})

	utils.AssertNil(err, t)
	utils.AssertEqual("1", data["x"], t)
	utils.AssertEqual("true", data["y"], t)
}

func TestParseDataWithArray(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": [1, 2, 3]}`)),
	})
	expectedSlice := []string{"1", "2", "3"}
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
