package crud

import (
	"api_meta/interfaces"
	"controllers/plugins/account_hash"
	"controllers/plugins/response_writer"
	"github.com/gorilla/mux"
	"net/http"
)

type controller struct {
	crudServicesPool interfaces.CRUDServicesPool
}

func Init(r *mux.Router, pool interfaces.CRUDServicesPool) {
	c := controller{crudServicesPool: pool}

	r.HandleFunc(
		"/{entity_type:[a-zA-Z-_]+?}/",
		c.getAll,
	).Methods(http.MethodGet)
	r.HandleFunc(
		"/{entity_type:[a-zA-Z-_]+?}/{entity_hash:[a-f0-9]{32}}/",
		c.getOne,
	).Methods(http.MethodGet)
	r.HandleFunc("/{entity_type:[a-zA-Z-_]+?}/", c.create).Methods(http.MethodPost)
	r.HandleFunc("/{entity_type:[a-zA-Z-_]+?}/", c.update).Methods(http.MethodPatch)
	r.HandleFunc(
		"/{entity_type:[a-zA-Z-_]+?}/{entity_hash:[a-f0-9]{32}}/",
		c.delete,
	).Methods(http.MethodDelete)
}

func (c controller) getAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType := vars["entity_type"]

	response := c.crudServicesPool.GetReadService(entityType).GetAll(account_hash.ExtractFromRequest(r))
	response_writer.Write(w, response)
}

func (c controller) getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType, entityHash := vars["entity_type"], vars["entity_hash"]

	response := c.crudServicesPool.GetReadService(entityType).Get(
		account_hash.ExtractFromRequest(r),
		entityHash,
	)
	response_writer.Write(w, response)
}

func (c controller) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType := vars["entity_type"]

	response := c.crudServicesPool.GetCreateService(entityType).Create(
		account_hash.ExtractFromRequest(r),
		r.Body,
	)
	response_writer.Write(w, response)
}

func (c controller) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType := vars["entity_type"]

	response := c.crudServicesPool.GetUpdateService(entityType).Update(
		account_hash.ExtractFromRequest(r),
		r.Body,
	)
	response_writer.Write(w, response)
}

func (c controller) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityType, entityHash := vars["entity_type"], vars["entity_hash"]

	response := c.crudServicesPool.GetDeleteService(entityType).Delete(
		account_hash.ExtractFromRequest(r),
		entityHash,
	)
	response_writer.Write(w, response)
}
