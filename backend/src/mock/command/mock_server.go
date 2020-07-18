package command

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"plugins/logger"
	"sync"
)

type (
	user struct {
		Hash string
		Name string
	}

	handler struct {
		users map[string]user
		*sync.Mutex
	}

	response struct {
		status string
	}

	successResponse struct {
		response
		data interface{}
	}

	erroredResponse struct {
		response
		errorDetail string
	}
)

func Init(router *mux.Router) {
	h := handler{users: map[string]user{}}

	router.HandleFunc("/", h.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/?{hash:[0-9]+}", h.getUser).Methods(http.MethodGet)
	router.HandleFunc("/", h.createUser).Methods(http.MethodPost)
	router.HandleFunc("/?{hash:[0-9]+}", h.patchUser).Methods(http.MethodPatch)
	router.HandleFunc("/?{hash:[0-9]+}", h.deleteUser).Methods(http.MethodDelete)
}

func (h handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	defer sendErrorIfPanicked(w)

	encodeAndSendResponse(w, h.users)
}

func (h handler) getUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
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
	h.Lock()
	defer h.Unlock()
	defer sendErrorIfPanicked(w)

	var u user
	decodeRequestBody(r, &u)

	_, userExists := h.users[u.Hash]
	if userExists {
		panic(errors.New("user-already-exists"))
	} else {
		encodeAndSendResponse(w, nil)
	}
}

func (h *handler) patchUser(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
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
	h.Lock()
	defer h.Unlock()
	defer sendErrorIfPanicked(w)

	delete(h.users, mux.Vars(r)["hash"])
	encodeAndSendResponse(w, nil)
}

func sendErrorIfPanicked(w http.ResponseWriter) {
	if err := recover(); err != nil {
		logger.WarningF("Panicked: %v", err)

		output, _ := json.Marshal(erroredResponse{
			response:    response{status: "error"},
			errorDetail: err.(error).Error(),
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
		response: response{status: "ok"},
		data:     v,
	})
	makeResponse(w, output)
}

func makeResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("content-type", "application/json")

	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}
