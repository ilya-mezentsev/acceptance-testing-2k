package tests_runner

import (
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
	"services/tests_runner/runner"
)

type controller struct {
	service runner.Service
}

func Init(
	r *mux.Router,
	service runner.Service,
	middlewares ...mux.MiddlewareFunc,
) {
	c := controller{service}
	testsRunnerAPI := r.PathPrefix("/tests").Subrouter()
	testsRunnerAPI.Use(middlewares...)

	testsRunnerAPI.HandleFunc("/{account_hash:[a-f0-9]{32}}", c.run).Methods(http.MethodPost)
}

func (c controller) run(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountHash := vars["account_hash"]

	response_writer.Write(w, c.service.Run(accountHash, r))
}
