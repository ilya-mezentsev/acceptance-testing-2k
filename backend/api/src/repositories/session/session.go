package session

import (
	"api_meta/interfaces"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.SessionRepository {
	return repository{db}
}

func (r repository) AccountExists(accountHash string) (bool, error) {
	var accountExists bool
	err := r.db.Get(&accountExists, `SELECT 1 FROM accounts WHERE hash = ?`, accountHash)
	if err == sql.ErrNoRows {
		accountExists = false
		err = nil
	}

	return accountExists, err
}
