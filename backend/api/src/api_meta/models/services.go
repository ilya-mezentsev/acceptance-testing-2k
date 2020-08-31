package models

type (
	UpdateRequest struct {
		AccountHash   string        `json:"account_hash" validation:"md5-hash"`
		UpdatePayload []UpdateModel `json:"update_payload"`
	}

	UpdateMetaRequest struct {
		AccountHash string        `json:"account_hash" validation:"md5-hash"`
		Headers     []UpdateModel `json:"headers"`
		Cookies     []UpdateModel `json:"cookies"`
	}
)
