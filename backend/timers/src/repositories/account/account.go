package account

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) GetNonVerifiedAccountsCreatedDayAgo() ([]string, error) {
	var hashes []string
	err := r.db.Select(
		&hashes,
		`SELECT hash FROM accounts WHERE created_at < datetime('now', '-1 day') AND verified = 0`,
	)

	return hashes, err
}

func (r Repository) DeleteAccounts(hashes []string) error {
	_, err := r.db.Exec(
		`DELETE FROM accounts WHERE hash IN (?)`,
		strings.Join(hashes, ","),
	)

	return err
}
