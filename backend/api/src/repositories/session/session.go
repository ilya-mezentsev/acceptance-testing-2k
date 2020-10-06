package session

import (
	"api_meta/interfaces"
	"database/sql"
	"db_connector"
)

type repository struct {
	connector *db_connector.Connector
}

func New(connector *db_connector.Connector) interfaces.SessionRepository {
	return repository{connector}
}

func (r repository) CredentialsExists(
	accountHash,
	login,
	password string,
) (bool, error) {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return false, err
	}

	var accountExists bool
	err = db.Get(
		&accountExists,
		`SELECT 1 FROM account_credentials WHERE login = ? AND password = ?`,
		login, password,
	)
	if err == sql.ErrNoRows {
		accountExists = false
		err = nil
	}

	return accountExists, err
}
