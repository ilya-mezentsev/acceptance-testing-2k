package tests_runner

import (
	"controllers/plugins/account_hash"
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
	"services/tests_runner/runner"
)

func Init(r *mux.Router, service runner.Service) {
	r.HandleFunc("/tests/", func(w http.ResponseWriter, r *http.Request) {
		response_writer.Write(w, service.Run(account_hash.ExtractFromRequest(r), r))
	}).Methods(http.MethodPost)
}
