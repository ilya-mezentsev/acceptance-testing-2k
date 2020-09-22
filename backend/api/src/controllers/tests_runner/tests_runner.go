package tests_runner

import (
	"controllers/plugins/account_hash"
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
	"services/tests_runner/file_creator"
	"services/tests_runner/runner"
)

func Init(
	r *mux.Router,
	fileCreator file_creator.Service,
	testsRunner runner.Service,
) {
	r.HandleFunc("/tests-file/", func(w http.ResponseWriter, r *http.Request) {
		response_writer.Write(w, fileCreator.CreateTestsFile(account_hash.ExtractFromRequest(r), r))
	}).Methods(http.MethodPost)

	r.HandleFunc("/run-tests/", func(w http.ResponseWriter, r *http.Request) {
		testsRunner.Run(
			account_hash.ExtractFromRequest(r),
			w,
			r,
		)
	}).Methods(http.MethodGet)
}
