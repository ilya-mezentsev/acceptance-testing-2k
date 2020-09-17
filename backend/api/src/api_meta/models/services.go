package models

type (
	UpdateRequest struct {
		UpdatePayload []UpdateModel `json:"update_payload"`
	}

	UpdateMetaRequest struct {
		Headers []UpdateModel `json:"headers"`
		Cookies []UpdateModel `json:"cookies"`
	}
)
