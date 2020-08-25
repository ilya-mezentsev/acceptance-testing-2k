package models

type SessionResponse struct {
	Login       string `json:"login"`
	AccountHash string `json:"account_hash"`
}
