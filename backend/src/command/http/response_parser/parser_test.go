package response_parser

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"utils"
)

func TestParseData(t *testing.T) {
	data, err := Parse(http.Response{
		Status:     "Ok",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"x": 1, "y": true}`)),
	})

	utils.AssertNil(err, t)
	utils.AssertEqual("1", data["x"], t)
	utils.AssertEqual("true", data["y"], t)
}
