package registration

import (
	"api_meta/interfaces"
	"api_meta/models"
	"api_meta/types"
	"db_connector"
	"github.com/jmoiron/sqlx"
	sqlite "github.com/mattn/go-sqlite3"
)

const (
	addAccountQuery            = `INSERT INTO accounts(hash) VALUES(?)`
	addAccountCredentialsQuery = `
	INSERT INTO account_credentials(login, password)
	VALUES(:login, :password)`
)

type repository struct {
	db        *sqlx.DB
	connector db_connector.Connector
}

func New(db *sqlx.DB, connector db_connector.Connector) interfaces.RegistrationRepository {
	return repository{db: db, connector: connector}
}

func (r repository) CreateAccount(accountHash string) error {
	_, err := r.db.Exec(addAccountQuery, accountHash)
	sqliteError, ok := err.(sqlite.Error)
	if ok && sqliteError.Code == sqlite.ErrConstraint {
		return types.AccountHashAlreadyExists{}
	}

	return err
}

func (r repository) CreateAccountCredentials(record models.AccountCredentialsRecord) error {
	db, err := r.connector.Connect(record.AccountHash)
	if err != nil {
		return err
	}

	_, err = db.NamedExec(addAccountCredentialsQuery, record)
	return err
}
