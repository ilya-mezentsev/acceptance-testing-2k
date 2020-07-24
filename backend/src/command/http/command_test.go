package http

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	mockCommand "mock/command"
	"net/http"
	"os"
	"testing"
	"utils"
)

var r = mux.NewRouter()

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestCommand_Run(t *testing.T) {
	mockCommand.Init(r)
	server := utils.GetTestServer(r)
	defer server.Close()

	res, err := New(mockCommand.MockSettings{
		Method:             http.MethodGet,
		BaseURL:            server.URL,
		Endpoint:           "users",
		PassArgumentsInURL: false,
	}).Run(``)

	utils.AssertNil(err, t)
	t.Log(res)
}
