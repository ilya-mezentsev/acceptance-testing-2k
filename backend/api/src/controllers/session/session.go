package session

import (
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
	"services/session"
)

type controller struct {
	service *session.Service
}

func Init(r *mux.Router, service *session.Service) {
	c := controller{service}
	sessionAPI := r.PathPrefix("/session").Subrouter()

	sessionAPI.HandleFunc("/", c.getSession).Methods(http.MethodGet)
	sessionAPI.HandleFunc("/", c.createSession).Methods(http.MethodPost)
	sessionAPI.HandleFunc("/", c.deleteSession).Methods(http.MethodDelete)
}

func (c controller) getSession(w http.ResponseWriter, r *http.Request) {
	response_writer.Write(w, c.service.GetSession(w, r))
}

func (c controller) createSession(w http.ResponseWriter, r *http.Request) {
	response_writer.Write(w, c.service.CreateSession(w, r))
}

func (c controller) deleteSession(w http.ResponseWriter, _ *http.Request) {
	response_writer.Write(w, c.service.DeleteSession(w))
}
