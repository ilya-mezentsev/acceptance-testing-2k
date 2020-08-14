package registration

import (
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
	"services/registration"
)

func Init(r *mux.Router, service registration.Service) {
	registrationAPI := r.PathPrefix("/registration").Subrouter()

	registrationAPI.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response_writer.Write(w, service.Register(r.Body))
	}).Methods(http.MethodPost)
}
