package response_writer

import (
	"api_meta/interfaces"
	"encoding/json"
	"logger"
	"net/http"
)

type (
	Response struct {
		Status string `json:"status"`
	}

	ResponseWithData struct {
		Response
		Data interface{} `json:"data"`
	}
)

func Write(w http.ResponseWriter, r interfaces.Response) {
	if r.HasData() {
		writeResponse(w, ResponseWithData{
			Response: Response{
				Status: r.GetStatus(),
			},
			Data: r.GetData(),
		})
	} else {
		writeResponse(w, Response{Status: r.GetStatus()})
	}
}

func writeResponse(w http.ResponseWriter, response interface{}) {
	data, _ := json.Marshal(response)

	if _, err := w.Write(data); err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to write response data: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"response": response,
			},
		}, logger.Error)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
