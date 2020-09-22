package runner

import (
	"github.com/gorilla/websocket"
	"logger"
	"net/http"
	"services/errors"
	"services/tests_runner/client"
	"services/tests_runner/plugins/tests_file_path"
)

type Service struct {
	client   client.Grpc
	upgrader websocket.Upgrader
}

func New(client client.Grpc) Service {
	return Service{
		client: client,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s Service) Run(
	accountHash string,
	w http.ResponseWriter,
	r *http.Request,
) {
	filename := r.URL.Query().Get("filename")
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ErrorF("Unable to upgrade connection: %v", err)
		return
	}
	var response interface{} = nil
	report, err := s.client.Call(
		accountHash,
		tests_file_path.BuildFilePath(accountHash, filename),
	)
	if err != nil {
		response = errors.ServiceError{
			Code:        unableToRunTestsCode,
			Description: callRemoteProcedureError,
		}
	} else {
		response = report
	}

	err = c.WriteJSON(response)
	defer func() {
		_ = c.Close()
	}()

	if err != nil {
		logger.WarningF("Unable to send response to connection: %v", err)
		return
	}
}
