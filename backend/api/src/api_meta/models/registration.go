package models

type (
	RegistrationRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	CreateSessionRequest RegistrationRequest

	AccountCredentialsRecord struct {
		AccountHash string `db:"account_hash"`
		Login       string `db:"login"`
		Password    string `db:"password"`
	}
)
