package command

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"plugins/logger"
)

type (
	user struct {
		Hash string
		Name string
	}

	handler struct {
		users map[string]user
	}

	Response struct {
		Status string
	}

	successResponse struct {
		Response
		Data interface{}
	}

	erroredResponse struct {
		Response
		ErrorDetail string
	}
)

func Init(router *mux.Router) {
	h := handler{users: map[string]user{
		"hash-1": {
			Hash: "hash-1",
			Name: "John",
		},
		"hash-2": {
			Hash: "hash-2",
			Name: "Nick",
		},
	}}
	userAPI := router.PathPrefix("/user").Subrouter()

	router.HandleFunc("/users", h.getAllUsers).Methods(http.MethodGet)
	userAPI.HandleFunc("/{hash:[a-zA-Z0-9-]+}", h.getUser).Methods(http.MethodGet)
	userAPI.HandleFunc("/", h.createUser).Methods(http.MethodPost)
	userAPI.HandleFunc("/{hash:[a-zA-Z0-9-]+}", h.patchUser).Methods(http.MethodPatch)
	userAPI.HandleFunc("/{hash:[a-zA-Z0-9-]+}", h.deleteUser).Methods(http.MethodDelete)
}

func (h handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	defer sendErrorIfPanicked(w)

	encodeAndSendResponse(w, h.users)
}

func (h handler) getUser(w http.ResponseWriter, r *http.Request) {
	defer sendErrorIfPanicked(w)

	userHash := mux.Vars(r)["hash"]
	user, found := h.users[userHash]
	if found {
		encodeAndSendResponse(w, user)
	} else {
		panic(errors.New("user-not-found"))
	}
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	defer sendErrorIfPanicked(w)

	var u user
	decodeRequestBody(r, &u)

	_, userExists := h.users[u.Hash]
	if userExists {
		panic(errors.New("user-already-exists"))
	} else {
		h.users[u.Hash] = u
		encodeAndSendResponse(w, nil)
	}
}

func (h *handler) patchUser(w http.ResponseWriter, r *http.Request) {
	defer sendErrorIfPanicked(w)

	var u user
	decodeRequestBody(r, &u)

	_, userExists := h.users[u.Hash]
	if userExists {
		h.users[u.Hash] = u
	} else {
		panic(errors.New("user-not-found"))
	}
}

func (h *handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	defer sendErrorIfPanicked(w)

	delete(h.users, mux.Vars(r)["hash"])
	encodeAndSendResponse(w, nil)
}

func sendErrorIfPanicked(w http.ResponseWriter) {
	if err := recover(); err != nil {
		logger.WarningF("Panicked: %v", err)

		output, _ := json.Marshal(erroredResponse{
			Response:    Response{Status: "error"},
			ErrorDetail: err.(error).Error(),
		})
		makeResponse(w, output)
	}
}

func decodeRequestBody(r *http.Request, target interface{}) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(requestBody, target)
	if err != nil {
		panic(err)
	}
}

func encodeAndSendResponse(w http.ResponseWriter, v interface{}) {
	output, _ := json.Marshal(successResponse{
		Response: Response{Status: "ok"},
		Data:     v,
	})
	makeResponse(w, output)
}

func makeResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("content-type", "application/json")

	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}
