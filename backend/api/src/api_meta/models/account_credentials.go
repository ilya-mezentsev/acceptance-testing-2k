package models

type (
	AccountCredentials struct {
		Hash     string `db:"hash"`
		Login    string `db:"login"`
		Verified bool   `db:"verified"`
	}
)
