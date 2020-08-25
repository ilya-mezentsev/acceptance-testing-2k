package crud

import (
	"api_meta/interfaces"
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
)

type controller struct {
	crudServicesPool interfaces.CRUDServicesPool
}

func Init(r *mux.Router, pool interfaces.CRUDServicesPool) {
	c := controller{crudServicesPool: pool}
	crudAPI := r.PathPrefix("/entity").Subrouter()

	crudAPI.HandleFunc("/{entity_type:[a-zA-Z-_]+?}/{account_hash:[a-f0-9]{32}}/", c.getAll).Methods(http.MethodGet)
	crudAPI.HandleFunc(
		"/{entity_type:[a-zA-Z-_]+?}/{account_hash:[a-f0-9]{32}}/{entity_hash:[a-f0-9]+?}/",
		c.getOne,
	).Methods(http.MethodGet)
	crudAPI.HandleFunc("/{entity_type:[a-zA-Z-_]+?}/", c.create).Methods(http.MethodPost)
	crudAPI.HandleFunc("/{entity_type:[a-zA-Z-_]+?}/", c.update).Methods(http.MethodPatch)
	crudAPI.HandleFunc(
		"/{entity_type:[a-z_]+?}/{account_hash:[a-f0-9]{32}}/{entity_hash:[a-f0-9]+?}/",
		c.delete,
	).Methods(http.MethodDelete)
}

func (c controller) getAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType, accountHash := vars["entity_type"], vars["account_hash"]

	response := c.crudServicesPool.Get(entityType).GetAll(accountHash)
	response_writer.Write(w, response)
}

func (c controller) getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType, accountHash, entityHash := vars["entity_type"], vars["account_hash"], vars["entity_hash"]

	response := c.crudServicesPool.Get(entityType).Get(accountHash, entityHash)
	response_writer.Write(w, response)
}

func (c controller) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType := vars["entity_type"]

	response := c.crudServicesPool.Get(entityType).Create(r.Body)
	response_writer.Write(w, response)
}

func (c controller) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType := vars["entity_type"]

	response := c.crudServicesPool.Get(entityType).Update(r.Body)
	response_writer.Write(w, response)
}

func (c controller) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType, accountHash, entityHash := vars["entity_type"], vars["account_hash"], vars["entity_hash"]

	response := c.crudServicesPool.Get(entityType).Delete(accountHash, entityHash)
	response_writer.Write(w, response)
}
