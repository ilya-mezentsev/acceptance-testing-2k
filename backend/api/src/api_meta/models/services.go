package models

type (
	UpdateRequest struct {
		AccountHash   string        `json:"account_hash" validation:"md5-hash"`
		UpdatePayload []UpdateModel `json:"update_payload"`
	}
)
